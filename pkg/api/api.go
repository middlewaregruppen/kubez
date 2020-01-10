package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type API struct {
	*Delay
	*TransmissionRate
	Name         string      `json:"name"`
	Port         int64       `json:"port"`
	Path         string      `json:"path"`
	CORS         bool        `json:"cors"`
	LogToConsole bool        `json:"logToConsole"`
	FailureRate  FailureRate `json:"failureRate"`

	Response Response `json:"response"`
	Self     string   `json:"self"`
}

type FailureRate struct {
	Rate          int `json:"rate"`
	ResponseCodes int `json:"responseCodes"`
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

	if a.TransmissionRate != nil {
		body = a.TransmissionRate.RequestReader(r)
	}

	if a.Delay != nil {
		a.DelayRequest()
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

	var bufSize int64
	bufSize = 512 * 1024 // Default buffert size

	if a.TransmissionRate != nil {
		response, bufSize = a.TransmissionRate.ResponseReader(response)
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
