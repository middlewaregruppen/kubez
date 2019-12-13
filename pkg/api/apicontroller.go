package api

import (
	"errors"
	"log"
)

type APIController struct {
	APICollection []*API `json:"apis"`
}

func (ac *APIController) RegisterAPI(api *API) {
	ac.APICollection = append(ac.APICollection, api)
}

func (ac *APIController) GetEndpoint(name string) *API {
	//Find API
	for _, api := range ac.APICollection {
		if api.Name == name {
			return api
		}
	}
	return nil
}

func (ac *APIController) CreateEndpoint(a *API) error {
	//Find API
	exists := ac.GetEndpoint(a.Name)
	if exists != nil {
		return errors.New("Endpoint exists")
	}
	ac.APICollection = append(ac.APICollection, a)
	return nil
}

func (ac APIController) GetEndpointList() []*API {
	return ac.APICollection
}

func (ac *APIController) UpdateEndpoint(a *API) error {
	//Find API
	for i, api := range ac.APICollection {
		if api.Name == a.Name {
			log.Print("Found it ")
			ac.APICollection[i] = a
			return nil
		}
	}
	return errors.New("NotFound")
}
