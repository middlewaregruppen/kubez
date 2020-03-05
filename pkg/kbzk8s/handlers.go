package kbzk8s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HandleLoad is a webhandler to create load.
func HandleLoad(rw http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("%s", err)
		rw.Write([]byte(msg))
		return
	}

	k8sreq := &K8SLoad{}
	err = json.Unmarshal(b, k8sreq)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		es := fmt.Sprintf("Error processing incomming data: %s", err)
		rw.Write([]byte(es))
		return
	}
	err = Load(k8sreq)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		es := fmt.Sprintf("Error processing request: %s", err)
		rw.Write([]byte(es))
		return
	}
}

// HandleGetPodList is a webhandler for getting all pods in cluster.
func HandleGetPodList(rw http.ResponseWriter, r *http.Request) {

	pil, err := GetPodInfoList()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		es := fmt.Sprintf("%s", err)
		rw.Write([]byte(es))
		return
	}

	b, err := json.Marshal(pil)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(b)
}
