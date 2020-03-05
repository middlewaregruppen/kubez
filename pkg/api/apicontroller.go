package api

import (
	"fmt"
	"log"
)

// Controller ...
type Controller struct {
}

// CreateEndpoint creates an new APIEdnpoint
func (ac *Controller) CreateEndpoint(a *API) error {

	deploymentName := fmt.Sprintf("kubezapi-%s", a.Name)

	err := upsertConfigMap(a.Namespace, deploymentName, a)

	if err != nil {
		return err
	}
	err = createAPIEndpointDeployment(a.Namespace, deploymentName, a)
	if err != nil {
		return err
	}

	err = createAPIEndpointService(a.Namespace, deploymentName, a)
	if err != nil {
		return err
	}

	return nil
}

// GetEndpointList returns a list of all Endpoints.
func (ac *Controller) GetEndpointList() ([]*API, error) {

	dlst, err := getAPIEndpointDeployments()
	if err != nil {
		return nil, err
	}

	var apis []*API

	// Iterate though deployments.
	for _, d := range dlst {

		//  Get the corresponding config map
		api, err := getAPIConfig(d.Namespace, d.Labels["kubez-name"])
		if err != nil {
			log.Printf("Unable to get apiconfig: %s", err)
			continue
		}
		api.Namespace = d.Namespace
		api.Status.RunningPods = d.Status.ReadyReplicas

		apis = append(apis, api)

	}
	return apis, nil

}

// UpdateEndpoint updates the endpoint
func (ac *Controller) UpdateEndpoint(a *API) error {
	deploymentName := fmt.Sprintf("kubezapi-%s", a.Name)

	return upsertConfigMap(a.Namespace, deploymentName, a)
}
