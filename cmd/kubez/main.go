package main

import (
	"log"
	"fmt"
	"time"
	"flag"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/middlewaregruppen/kubez/pkg/actions"
	"github.com/middlewaregruppen/kubez/pkg/api"
	"github.com/middlewaregruppen/kubez/pkg/info"
	"github.com/middlewaregruppen/kubez/pkg/kbzk8s"
	"github.com/middlewaregruppen/kubez/pkg/network"
)

// VERSION is generated during compile as is never to be set here
var VERSION string

// COMMIT is the Git commit hash and is generated during compile as is never to be set here
var COMMIT string

// BRANCH is the Git branch name and is generated during compile as is never to be set here
var BRANCH string

// GOVERSION is the Go version used to compile and is generated during compile as is never to be set here
var GOVERSION string

func main() {

	// Request app version
	showver := flag.CommandLine.Bool("version", false, "Print version")
	flag.Parse()

	// Show version if requested
	if *showver {
		fmt.Printf("Version: %s\nCommit: %s\nBranch: %s\nGoVersion: %s\n", VERSION, COMMIT, BRANCH, GOVERSION)
		return
	}

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
