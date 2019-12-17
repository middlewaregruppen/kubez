package info

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//GetHTTPHeaders - get http headers
func GetHTTPHeaders(r *http.Request) []byte {
	var out []string

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			out = append(out, fmt.Sprintf("%v:%v", name, h))
		}
	}
	outJSON, _ := json.Marshal(out)

	return outJSON
}
