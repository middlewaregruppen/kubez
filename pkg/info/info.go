package info

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	CGroup *CGroup `json:"cGroup"`
}

func HandleGetInfo(w http.ResponseWriter, r *http.Request) {

	cg := GetCgroup()
	rt := &Info{
		CGroup: cg,
	}

	b, _ := json.Marshal(rt)
	w.Header().Add("Content-Type:", "application/json")
	w.Write(b)
}
