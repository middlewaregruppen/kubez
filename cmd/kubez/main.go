package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/middlewaregruppen/kubez/pkg/actions"
	"github.com/middlewaregruppen/kubez/pkg/api"
	"github.com/middlewaregruppen/kubez/pkg/info"
)

func main() {

	log.Printf("Starting kubez")
	//pkger.Include("/web/frontend/dist")

	ac := &api.APIController{}

	ac.RegisterAPI(&api.API{
		Name: "testa",
		Path: "/api/hello",
		Port: 4000,
		Response: api.Response{
			Type:   "static",
			Static: "world",
		},
	})
	ac.RegisterAPI(&api.API{
		Name: "Test B",
		Path: "/api/hello2",
		Port: 4001,
		Response: api.Response{
			Type:   "static",
			Static: "worldss",
		},
	})

	r := mux.NewRouter()
	r.HandleFunc("/kubez/info", info.HandleGetInfo).Methods("GET")
	r.HandleFunc("/kubez/action/{action}", actions.ActionHandler).Methods("POST")
	//r.PathPrefix("/api/").Handler(ac)

	r.HandleFunc("/kubez/apicc/", ac.HandleGetEndpointList).Methods("GET")
	r.HandleFunc("/kubez/apicc/{endpoint}", ac.HandleUpdateEndpoint).Methods("PUT")
	r.HandleFunc("/kubez/apicc/{endpoint}", ac.HandleGetEndpoint).Methods("GET")
	r.HandleFunc("/kubez/apicc/", ac.HandleCreateEndpoint).Methods("POST")

	dir := http.FileServer(http.Dir("web/frontend/dist"))
	r.PathPrefix("/").Handler(dir)

	http.ListenAndServe(":3000", r)
}
