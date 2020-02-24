package info

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/middlewaregruppen/kubez/pkg/kbzk8s"
)

type Info struct {
	CGroup   *CGroup             `json:"cGroup"`
	Headers  map[string][]string `json:"httpheaders"`
	Hostname string              `json:"hostname"`
	K8SStats *kbzk8s.Stats       `json:"k8sstats"`
}

//HandleGetInfo - get information
func HandleGetInfo(w http.ResponseWriter, r *http.Request) {

	cg := GetCgroup()
	httpheaders := GetHTTPHeaders(r)

	ki := kbzk8s.GetStats()

	rt := &Info{
		CGroup:   cg,
		Headers:  httpheaders,
		K8SStats: ki,
	}
	hn, _ := os.Hostname()
	rt.Hostname = hn

	b, _ := json.Marshal(rt)
	w.Header().Add("Content-Type:", "application/json")
	w.Write(b)
}
