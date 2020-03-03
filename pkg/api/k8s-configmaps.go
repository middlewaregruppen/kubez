package api

import (
	"encoding/json"
	"errors"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// upsertConfigMap creates a new config map if it does not exists. If it exists it will be updated.
func upsertConfigMap(namespace, name string, api *API) error {

	data, err := json.Marshal(api)
	if err != nil {
		return err
	}

	cm := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"kubez-type": "api-endpoint",
				"kubez-name": api.Name,
			},
		},
		Data: map[string]string{
			"api.json": fmt.Sprintf("%s", data),
		},
	}

	clientset, err := getClientSet()
	if err != nil {
		return err
	}

	// Try to create
	_, err = clientset.CoreV1().ConfigMaps(namespace).Create(cm)

	// Successful?
	if err == nil {
		return nil
	}

	// Create failed, try to update.
	_, err = clientset.CoreV1().ConfigMaps(namespace).Update(cm)

	return err
}

// getAPIConfig returns the config as an API struct pointer from the configmap
func getAPIConfig(namespace, apiname string) (*API, error) {

	// creates the clientset
	clientset, err := getClientSet()

	if err != nil {
		return nil, err
	}
	selector := fmt.Sprintf("kubez-type=api-endpoint, kubez-name=%s", apiname)

	cms, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{
		LabelSelector: selector,
	})

	if err != nil {
		return nil, err
	}
	if len(cms.Items) == 0 {
		return nil, errors.New("Not found")
	}

	data := cms.Items[0].Data["api.json"]

	api := &API{}
	err = json.Unmarshal([]byte(data), api)
	if err != nil {
		return nil, err
	}
	api.Namespace = namespace

	return api, nil

}
