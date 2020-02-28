package api

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//createAPIEndpointService creates an new service for an api-endpoint.
func createAPIEndpointService(namespace, serviceName string, api *API) error {

	clientset, err := getClientSet()

	if err != nil {
		return err
	}

	var stype v1.ServiceType

	switch api.ServiceType {

	case "clusterIP":
		stype = v1.ServiceTypeClusterIP

	case "nodePort":
		stype = v1.ServiceTypeNodePort

	case "loadBalancer":
		stype = v1.ServiceTypeLoadBalancer

	default:
		stype = v1.ServiceTypeClusterIP

	}

	_, err = clientset.CoreV1().Services(namespace).Create(&v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
			Labels: map[string]string{
				"kubez-type": "api-endpoint",
				"kubez-name": api.Name,
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"kubez-type": "api-endpoint",
				"kubez-name": api.Name,
			},
			Type: stype,
			Ports: []v1.ServicePort{
				{
					Name:       "kubez-api",
					Port:       api.Port,
					TargetPort: intstr.IntOrString{IntVal: api.Port},
				},
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
