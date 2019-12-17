package info

import (
	"encoding/json"
	"log"
	"net/http"
)

type Info struct {
	CGroup  *CGroup `json:"cGroup"`
	Headers []byte  `json:"headers"`
}

func HandleGetInfo(w http.ResponseWriter, r *http.Request) {

	cg := GetCgroup()
	headers := GetHTTPHeaders(r)

	rt := &Info{
		CGroup:  cg,
		Headers: headers,
	}

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
	b, _ := json.Marshal(rt)
	log.Println("DATA:", string(b))
	w.Header().Add("Content-Type:", "application/json")
	w.Write(b)
}
