package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// HandleGetEndpointList web handler for retriving an list of all available Endpoints
func (ac *Controller) HandleGetEndpointList(w http.ResponseWriter, r *http.Request) {
	eps, err := ac.GetEndpointList()

	if err != nil {
		log.Printf("Error getting APIEndpointlist %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	j, _ := json.Marshal(eps)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HandleCreateEndpoint web handler for creating an endpoint
func (ac *Controller) HandleCreateEndpoint(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)

	a := &API{}
	err := json.Unmarshal(b, a)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ac.CreateEndpoint(a)
	if err != nil {
		log.Printf("Error creating API Endpoint: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
}

//HandleUpdateEndpoint webhandler for updating an Endpoint
func (ac *Controller) HandleUpdateEndpoint(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)

	a := &API{}
	err := json.Unmarshal(b, a)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ac.UpdateEndpoint(a)

}
