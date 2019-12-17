package info

import (
	"net/http"
)

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
