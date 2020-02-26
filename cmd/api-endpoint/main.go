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

func main() {

	var config string
	flag.CommandLine.StringVar(&config, "config", "/config/api.json", "Config file to use for the API endpoint")
	flag.Parse()

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
		time.Sleep(time.Second)
	}
}
