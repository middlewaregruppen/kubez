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

type API struct {
	Name             string           `json:"name"`
	Port             int64            `json:"port"`
	Path             string           `json:"path"`
	CORS             bool             `json:"cors"`
	LogToConsole     bool             `json:"logToConsole"`
	Delay            Delay            `json:"delay"`
	FailureRate      FailureRate      `json:"failureRate"`
	TransmissionRate TransmissionRate `json:"transmissionRate"`

	Response Response `json:"response"`
	Self     string   `json:"self"`
}

type Delay struct {
	MinTime int `json:"minTime"`
	MaxTime int `json:"maxTime"`
}

type FailureRate struct {
	Rate          int `json:"rate"`
	ResponseCodes int `json:"responseCodes"`
}

type TransmissionRate struct {
	TX int64 `json:"tx"`
	RX int64 `json:"rx"`
}

type Response struct {
	Type    string            `json:"type"`
	Static  string            `json:"static"`
	Headers map[string]string `json:"headers"`
}

type ApiLog struct {
	APIName     string `json:"apiName"`
	ResposeType string `json:"respType"`
	Protocol    string `json:"proto"`
	Method      string `json:"method"`
	RequestURI  string `json:"requestUri"`
	StatusCode  int    `json:"statusCode"`
	UserAgent   string `json:"userAgent"`
	RemoteAddr  string `json:"remoteAddr"`
}

func (a *API) HandleAPIRequest(w http.ResponseWriter, r *http.Request) {

	var body io.Reader

	lm := &ApiLog{
		APIName:     a.Name,
		ResposeType: a.Response.Type,
		Protocol:    r.Proto,
		Method:      r.Method,
		RequestURI:  r.RequestURI,
		UserAgent:   r.UserAgent(),
		RemoteAddr:  r.RemoteAddr,
	}

	// RX Rate limiting in place?
	if a.TransmissionRate.RX > 0 {
		bucket := ratelimit.NewBucketWithRate(float64(a.TransmissionRate.TX), a.TransmissionRate.TX)
		body = ratelimit.Reader(r.Body, bucket)
	} else {
		body = r.Body
	}

	// Calulate how long time to delay.
	delay := 0
	if a.Delay.MinTime < a.Delay.MaxTime {
		delay = rand.Intn(a.Delay.MaxTime-a.Delay.MinTime) + a.Delay.MinTime
	}

	// Delay
	if delay > 0 {
		d, err := time.ParseDuration(fmt.Sprintf("%dms", delay))
		if err != nil {
			log.Print(err)
		}
		time.Sleep(d)
	}

	// Should we fail?
	if a.FailureRate.Rate > 0 {
		p := rand.Intn(100)
		if p < a.FailureRate.Rate {
			w.WriteHeader(a.FailureRate.ResponseCodes)
			lm.StatusCode = a.FailureRate.ResponseCodes
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
	for h, v := range a.Response.Headers {
		w.Header().Add(h, v)
	}

	var response io.Reader

	switch a.Response.Type {

	case "static":
		response = strings.NewReader(a.Response.Static)

	case "echo":
		b, _ := ioutil.ReadAll(body)
		response = bytes.NewReader(b)
	}

	var out io.Reader
	var bufSize int64

	// TX Rate limiting in place?
	if a.TransmissionRate.TX > 0 {
		bucket := ratelimit.NewBucketWithRate(float64(a.TransmissionRate.TX), a.TransmissionRate.TX)
		out = ratelimit.Reader(response, bucket)
		bufSize = a.TransmissionRate.TX
		// Limit how large the buffer can be (0,5 mb)
		if bufSize > 512*1024 {
			bufSize = 512 * 1024
		}
	} else {
		out = response
		bufSize = 512 * 1024
	}

	buf := make([]byte, bufSize)

	for {
		n, err := out.Read(buf)
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
