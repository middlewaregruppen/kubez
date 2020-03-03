package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/juju/ratelimit"
)

// API describes an API endpoint
type API struct {
	FailureRate int `json:"failurerate,string"`
	FailureCode int `json:"failurecode,string"`

	ResponseType  string            `json:"responsetype"`
	StaticContent string            `json:"staticcontent"`
	Headers       map[string]string `json:"headers"`

	DelayMin int `json:"mindelay,string"`
	DelayMax int `json:"maxdelay,string"`

	ResponseRate int64 `json:"responserate,string"`
	RequestRate  int64 `json:"requestrate,string"`

	ServiceType  string `json:"servicetype"`
	Status       Status `json:"status"`
	Namespace    string `json:"namespace"`
	Name         string `json:"name"`
	Port         int32  `json:"port,string"`
	Path         string `json:"path"`
	CORS         bool   `json:"cors"`
	LogToConsole bool   `json:"logtoconsole"`
}

// Status shows the status of the deployment
type Status struct {
	RunningPods       int32  `json:"runningpods,string"`
	ServiceType       string `json:"servicetype"`
	ConnectionDetails string `json:"connectiondetails"`
}

// Response configures which type of response that shall be sent back to the client.

// ApiLog describes the JSON format of the logs sent to std out
type ApiLog struct {
	APIName     string `json:"apiName"`
	ResposeType string `json:"respType"`
	Protocol    string `json:"proto"`
	Method      string `json:"method"`
	RequestURI  string `json:"requestUri"`
	StatusCode  int    `json:"statusCode,string"`
	UserAgent   string `json:"userAgent"`
	RemoteAddr  string `json:"remoteAddr"`
}

func (a *API) RequestReader(r *http.Request) io.Reader {
	if a.RequestRate == 0 {
		return r.Body
	}

	bucket := ratelimit.NewBucketWithRate(float64(a.RequestRate), a.RequestRate)
	return ratelimit.Reader(r.Body, bucket)

}

func (a *API) ResponseReader(response io.Reader) (reader io.Reader, buffSize int64) {

	bucket := ratelimit.NewBucketWithRate(float64(a.ResponseRate), a.ResponseRate)
	return ratelimit.Reader(response, bucket), a.ResponseRate

}

func (a *API) DelayRequest() {
	// Calulate how long time to delay.
	delay := 0
	if a.DelayMin < a.DelayMax {
		delay = rand.Intn(a.DelayMax-a.DelayMin) + a.DelayMin
	}

	// Delay
	if delay > 0 {
		d, err := time.ParseDuration(fmt.Sprintf("%dms", delay))
		if err != nil {
			log.Print(err)
		}
		time.Sleep(d)
	}

}

// HandleAPIRequest is a web handler that processes requests
func (a *API) HandleAPIRequest(w http.ResponseWriter, r *http.Request) {

	var body io.Reader

	lm := &ApiLog{
		APIName:     a.Name,
		ResposeType: a.ResponseType,
		Protocol:    r.Proto,
		Method:      r.Method,
		RequestURI:  r.RequestURI,
		UserAgent:   r.UserAgent(),
		RemoteAddr:  r.RemoteAddr,
	}

	if a.RequestRate > 0 {
		body = a.RequestReader(r)
	}

	a.DelayRequest()

	// Should we fail?
	if a.FailureRate > 0 {
		p := rand.Intn(100)
		if p < a.FailureRate {
			w.WriteHeader(a.FailureCode)
			lm.StatusCode = a.FailureCode
			a.logRequest(lm)
			return
		}
	}

	// Is it a request to get CORS?
	if r.Method == http.MethodOptions && a.CORS {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PATCH, PUT, POST, GET, OPTIONS, DELETE, HEAD, TRACE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Add headers
	for h, v := range a.Headers {
		w.Header().Add(h, v)
	}

	var response io.Reader

	switch a.ResponseType {

	case "static":
		response = strings.NewReader(a.StaticContent)

	case "echo":
		b, _ := ioutil.ReadAll(body)
		response = bytes.NewReader(b)
	default:
		response = bytes.NewReader([]byte("Configuration error: Invalid response type defined!"))
	}

	var bufSize int64
	bufSize = 512 * 1024 // Default buffert size

	if a.ResponseRate > 0 {
		response, bufSize = a.ResponseReader(response)
	}

	buf := make([]byte, bufSize)

	for {
		n, err := response.Read(buf)
		if err == io.EOF {
			break
		}
		w.Write(buf[:n])

		// Flush if implemeentation supports it
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
	a.logRequest(lm)
}

func (a *API) logRequest(lm *ApiLog) {

	j, _ := json.Marshal(lm)

	if a.LogToConsole {
		log.Printf("%s", j)
	}

}
