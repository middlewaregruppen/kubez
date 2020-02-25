package info

import (
	"net/http"
)

type RequestInfo struct {
	Proto string `json:"proto"`
	Major int    `json:"major"`
	Minor int    `json:"minor"`
}

func GetHttpRequest(r *http.Request) RequestInfo {
	ri := RequestInfo{
		Proto: r.Proto,
		Major: r.ProtoMajor,
		Minor: r.ProtoMinor,
	}
	return ri

}

//GetHTTPHeaders - get http headers
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
