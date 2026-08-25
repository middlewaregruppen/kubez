package info

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/middlewaregruppen/kubez/pkg/kbzk8s"
)

// Info ...
type Info struct {
	CGroup      *CGroup             `json:"cGroup"`
	Headers     map[string][]string `json:"httpheaders"`
	Hostname    string              `json:"hostname"`
	K8SStats    *kbzk8s.Stats       `json:"k8sstats"`
	RequestInfo RequestInfo         `json:"requestInfo"`
}

// HandleGetInfo - get information
func HandleGetInfo(w http.ResponseWriter, r *http.Request) {

	cg := GetCgroup()
	httpheaders := GetHTTPHeaders(r)

	ki := kbzk8s.GetStats()
	ri := GetHTTPRequest(r)

	rt := &Info{
		CGroup:      cg,
		Headers:     httpheaders,
		K8SStats:    ki,
		RequestInfo: ri,
	}
	hn, err := os.Hostname()
	if err != nil {
		log.Printf("Failed to get hostname: %s", err)
	}
	rt.Hostname = hn

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(rt); err != nil {
		log.Printf("Failed to write info response: %s", err)
	}
}
