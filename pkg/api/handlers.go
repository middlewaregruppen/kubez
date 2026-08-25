package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// HandleGetEndpointList web handler for retrieving a list of all available endpoints.
func (ac *Controller) HandleGetEndpointList(w http.ResponseWriter, r *http.Request) {
	eps, err := ac.GetEndpointList()

	if err != nil {
		log.Printf("Error getting APIEndpointlist %s", err)
		http.Error(w, "Failed to get API endpoints", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(eps); err != nil {
		log.Printf("Failed to write API endpoint list: %s", err)
	}
}

// HandleCreateEndpoint web handler for creating an endpoint
func (ac *Controller) HandleCreateEndpoint(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	a := &API{}
	err = json.Unmarshal(b, a)

	if err != nil {
		http.Error(w, "Invalid API endpoint", http.StatusBadRequest)
		return
	}

	err = ac.CreateEndpoint(a)
	if err != nil {
		log.Printf("Error creating API Endpoint: %s", err)
		http.Error(w, "Failed to create API endpoint", http.StatusInternalServerError)
		return

	}
}

// HandleUpdateEndpoint webhandler for updating an Endpoint
func (ac *Controller) HandleUpdateEndpoint(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	a := &API{}
	err = json.Unmarshal(b, a)

	if err != nil {
		http.Error(w, "Invalid API endpoint", http.StatusBadRequest)
		return
	}
	if err := ac.UpdateEndpoint(a); err != nil {
		log.Printf("Error updating API Endpoint: %s", err)
		http.Error(w, "Failed to update API endpoint", http.StatusInternalServerError)
	}

}
