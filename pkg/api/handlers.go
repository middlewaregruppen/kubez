package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (ac *APIController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	to := ac.apisServing(r.URL.RequestURI())

	if len(to) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Api is not registered"))
		return
	}

	// Let the first one handle the request.
	to[0].HandleAPIRequest(w, r)

}

func (ac *APIController) HandleGetEndpointList(w http.ResponseWriter, r *http.Request) {
	eps := ac.GetEndpointList()

	// Set self
	for _, e := range eps {
		e.Self = fmt.Sprintf("%s/%s", strings.TrimSuffix(r.RequestURI, "/"), e.Name)
	}

	j, _ := json.Marshal(eps)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (ac *APIController) HandleGetEndpoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ep := ac.GetEndpoint(vars["endpoint"])
	if ep == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	j, _ := json.Marshal(ep)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (ac *APIController) HandleCreateEndpoint(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)

	a := &API{}
	err := json.Unmarshal(b, a)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ac.CreateEndpoint(a)
}

func (ac *APIController) HandleUpdateEndpoint(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)

	a := &API{}
	err := json.Unmarshal(b, a)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ac.UpdateEndpoint(a)

	// Write out the saved
	ac.HandleGetEndpoint(w, r)
}

func (ac *APIController) apisServing(path string) []*API {

	var res []*API
	//Find API
	for _, api := range ac.APICollection {
		if api.Path == path {
			res = append(res, api)
		}
	}
	return res
}
