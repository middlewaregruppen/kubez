package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/middlewaregruppen/kubez/pkg/actions"
	"github.com/middlewaregruppen/kubez/pkg/api"
	"github.com/middlewaregruppen/kubez/pkg/info"
	"github.com/middlewaregruppen/kubez/pkg/kbzk8s"
	"github.com/middlewaregruppen/kubez/pkg/network"
)

func main() {

	log.Printf("Starting kubez")

	rand.Seed(time.Now().UnixNano())

	r := mux.NewRouter()
	r.HandleFunc("/kubez/info", info.HandleGetInfo).Methods("GET")
	r.HandleFunc("/kubez/action/{action}", actions.ActionHandler).Methods("POST")

	ac := &api.Controller{}
	r.HandleFunc("/kubez/apicc/", ac.HandleGetEndpointList).Methods("GET")
	r.HandleFunc("/kubez/apicc/{endpoint}", ac.HandleUpdateEndpoint).Methods("PUT")
	//r.HandleFunc("/kubez/apicc/{endpoint}", ac.HandleGetEndpoint).Methods("GET")
	r.HandleFunc("/kubez/apicc/", ac.HandleCreateEndpoint).Methods("POST")

	// Connection Checker
	r.HandleFunc("/kubez/connectioncheck/{target}", network.HandleCheckConnection).Methods("GET")

	// DNS Lookup
	r.HandleFunc("/kubez/dnslookup/{type}/{name}", network.HandleDNSLookup).Methods("GET")

	// K8s Loader
	r.HandleFunc("/kubez/k8sload", kbzk8s.HandleLoad).Methods("POST")

	// PodInfoList
	r.HandleFunc("/kubez/kbzk8s/podlist", kbzk8s.HandleGetPodList).Methods("GET")

	dir := http.FileServer(http.Dir("web/frontend/dist"))
	r.PathPrefix("/").Handler(dir)

	http.ListenAndServe(":3000", r)
}
