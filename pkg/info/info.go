package info

import (
	"encoding/json"
	"net/http"
	"os"
)

type Info struct {
	CGroup   *CGroup             `json:"cGroup"`
	Headers  map[string][]string `json:"httpheaders"`
	Hostname string              `json:"hostname"`
}

//HandleGetInfo - get information
func HandleGetInfo(w http.ResponseWriter, r *http.Request) {

	cg := GetCgroup()
	httpheaders := GetHTTPHeaders(r)

	rt := &Info{
		CGroup:  cg,
		Headers: httpheaders,
	}

	/*
		var v interface{}
		json.Unmarshal(headers, &v)
		out := v.([]interface{})

		for k, v := range out {
			switch v := v.(type) {
			case string:
				log.Println(k, v, "(string)")
			default:
				log.Println(k, v, "(unknown)")
			}
		}
	*/
	hn, _ := os.Hostname()
	rt.Hostname = hn

	b, _ := json.Marshal(rt)
	w.Header().Add("Content-Type:", "application/json")
	w.Write(b)
}
