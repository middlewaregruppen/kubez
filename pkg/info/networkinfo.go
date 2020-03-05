package info

import (
	"net/http"
)

// RequestInfo ...
type RequestInfo struct {
	Proto string `json:"proto"`
	Major int    `json:"major"`
	Minor int    `json:"minor"`
}

// GetHTTPRequest ...
func GetHTTPRequest(r *http.Request) RequestInfo {
	ri := RequestInfo{
		Proto: r.Proto,
		Major: r.ProtoMajor,
		Minor: r.ProtoMinor,
	}
	return ri

}

// GetHTTPHeaders gets http headers
func GetHTTPHeaders(r *http.Request) http.Header {
	/*
		var out []string

		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, h := range headers {
				out = append(out, fmt.Sprintf("%v:%v", name, h))
			}
		}
		outJSON, _ := json.Marshal(out)

		return outJSON
	*/
	return r.Header
}
