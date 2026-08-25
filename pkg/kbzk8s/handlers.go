package kbzk8s

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// HandleLoad is a webhandler to create load.
func HandleLoad(rw http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	k8sreq := &K8SLoad{}
	err = json.Unmarshal(b, k8sreq)
	if err != nil {
		http.Error(rw, "Error processing incoming data: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = Load(k8sreq)
	if err != nil {
		http.Error(rw, "Error processing request: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// HandleGetPodList is a webhandler for getting all pods in cluster.
func HandleGetPodList(rw http.ResponseWriter, r *http.Request) {

	pil, err := GetPodInfoList()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(pil); err != nil {
		log.Printf("Failed to write pod list response: %s", err)
	}
}
