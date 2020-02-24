package kbzk8s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HandleLoad(rw http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Printf("%s", err)
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
