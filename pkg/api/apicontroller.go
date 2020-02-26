package api

import (
	"fmt"
	"log"
	"net/http"
)

type APIController struct {
	APICollection []*API `json:"apis"`
	servers       []*http.Server
}

func (ac *APIController) RegisterAPI(api *API) {

}

func (ac *APIController) GetEndpoint(name string) (*API, error) {
	ns, err := thisNamespace()
	if err != nil {
		return nil, err
	}
	return getAPIConfig(ns, name)
}

func (ac *APIController) CreateEndpoint(a *API) error {

	ns, err := thisNamespace()
	if err != nil {
		return err
	}
	deploymentName := fmt.Sprintf("kubezapi-%s", a.Name)
	err = upsertConfigMap(ns, deploymentName, a)
	log.Print(err)
	err = deployAPIEndpoint(ns, deploymentName, a)
	log.Print(err)

	return nil
}

func (ac APIController) GetEndpointList() []*API {
	ns, _ := thisNamespace()
	apis, _ := getAPIList(ns)
	return apis
}

func (ac *APIController) UpdateEndpoint(a *API) error {
	ns, _ := thisNamespace()
	deploymentName := fmt.Sprintf("kubezapi-%s", a.Name)

	return upsertConfigMap(ns, deploymentName, a)
}
