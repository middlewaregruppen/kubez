package api

import (
	"errors"
	"fmt"
	"net/http"
)

type APIController struct {
	APICollection []*API `json:"apis"`
	servers       []*http.Server
}

func (ac *APIController) RegisterAPI(api *API) {
	ac.CreateAPIServer(api.Port)
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

func (ac *APIController) CreateAPIServer(port int64) {
	addr := fmt.Sprintf(":%d", port)
	for _, s := range ac.servers {
		// Already Exists
		if s.Addr == addr {
			return
		}
	}
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: ac,
	}
	go s.ListenAndServe()
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
			ac.APICollection[i] = a
			if a.Port != api.Port {
				ac.CreateAPIServer(a.Port)
			}
			return nil
		}
	}

	return errors.New("NotFound")
}
