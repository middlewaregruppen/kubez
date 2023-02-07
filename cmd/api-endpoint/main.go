package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/middlewaregruppen/kubez/pkg/api"
)

var a *api.API

// VERSION is generated during compile as is never to be set here
var VERSION string

// COMMIT is the Git commit hash and is generated during compile as is never to be set here
var COMMIT string

// BRANCH is the Git branch name and is generated during compile as is never to be set here
var BRANCH string

// GOVERSION is the Go version used to compile and is generated during compile as is never to be set here
var GOVERSION string

func main() {

	var config string
	flag.CommandLine.StringVar(&config, "config", "/config/api.json", "Config file to use for the API endpoint")
	showver := flag.CommandLine.Bool("version", false, "Print version")
	flag.Parse()

	// Show version if requested
	if *showver {
		fmt.Printf("Version: %s\nCommit: %s\nBranch: %s\nGoVersion: %s\n", VERSION, COMMIT, BRANCH, GOVERSION)
		return
	}

	a = &api.API{}
	err := loadAPIConfig(config)

	if err != nil {
		panic(err)
	}

	go updateAPIConfig(config)

	http.HandleFunc("/", a.HandleAPIRequest)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.Port), nil))

}

func loadAPIConfig(config string) error {
	cfg, err := ioutil.ReadFile(config)

	if err != nil {
		return err
	}

	err = json.Unmarshal(cfg, a)

	if err != nil {
		return err
	}
	return nil
}

func updateAPIConfig(config string) {
	for {
		loadAPIConfig(config)
		time.Sleep(time.Second * 2)
	}
}
