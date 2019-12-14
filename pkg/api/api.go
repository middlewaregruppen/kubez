package api

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type API struct {
	// Name of the API
	Name string `json:"name"`
	// Path
	Path        string      `json:"path"`
	CORS        bool        `json:"cors"`
	Delay       Delay       `json:"delay"`
	FailureRate FailureRate `json:"failureRate"`
	Response    Response    `json:"response"`
	Self        string      `json:"self"`
}

type Delay struct {
	MinTime int `json:"minTime"`
	MaxTime int `json:"maxTime"`
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

func (a *API) HandleAPIRequest(w http.ResponseWriter, r *http.Request) {

	// Calulate how long time to delay.
	delay := 0
	if a.Delay.MinTime < a.Delay.MaxTime {
		delay = rand.Intn(a.Delay.MaxTime-a.Delay.MinTime) + a.Delay.MinTime
		log.Print(delay)
	}

	// Wait for the minimum delay time
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

	// Write headers
	for h, v := range a.Response.Headers {
		w.Header().Add(h, v)
	}

	w.Write([]byte(a.Response.Static))

}
