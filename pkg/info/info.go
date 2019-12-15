package info

import (
	"encoding/json"
	"net/http"
	"os"
)

type Info struct {
	CGroup   *CGroup `json:"cGroup"`
	Hostname string  `json:"hostname"`
}

func HandleGetInfo(w http.ResponseWriter, r *http.Request) {

	cg := GetCgroup()
	rt := &Info{
		CGroup: cg,
	}

	hn, _ := os.Hostname()
	rt.Hostname = hn

	b, _ := json.Marshal(rt)
	w.Header().Add("Content-Type:", "application/json")
	w.Write(b)
}
