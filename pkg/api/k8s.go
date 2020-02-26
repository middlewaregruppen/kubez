package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func deployAPIEndpoint(namespace, deploymentName string, api *API) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	var replicas int32
	replicas = 1

	_, err = clientset.AppsV1().Deployments(namespace).Create(&appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
			Labels: map[string]string{
				"kubez-type": "api-endpoint",
				"kubez-name": api.Name,
			},
		},
		Spec: appv1.DeploymentSpec{

			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentName,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":        deploymentName,
						"kubez-type": "api-endpoint",
						"kubez-name": api.Name,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{

						{
							Name:  "api-endpoint",
							Image: "docker.io/middlewaregruppen/kubez:dev", //CHANGEME!
							Ports: []v1.ContainerPort{
								{ContainerPort: api.Port},
							},
							Command: []string{"/api-endpoint"},
							Args:    []string{"-config", "/config/api.json"},
							VolumeMounts: []v1.VolumeMount{
								{
									MountPath: "/config/",
									Name:      "api-config",
								},
							},
						},
					},
					Volumes: []v1.Volume{
						{
							Name: "api-config",
							VolumeSource: v1.VolumeSource{
								ConfigMap: &v1.ConfigMapVolumeSource{
									LocalObjectReference: v1.LocalObjectReference{
										Name: deploymentName,
									},
								},
							},
						},
					},
				},
			},
		},
	})
	return err
}

func upsertConfigMap(namespace, name string, api *API) error {

	data, err := json.Marshal(api)
	if err != nil {
		return err
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
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

	// Try to create
	_, err = clientset.CoreV1().ConfigMaps(namespace).Create(cm)

	// Done!
	if err == nil {
		return nil
	}

	// Create failed, try to update.
	_, err = clientset.CoreV1().ConfigMaps(namespace).Update(cm)

	return err
}

func getAPIList(namespace string) ([]*API, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	cms, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{
		LabelSelector: "kubez-type=api-endpoint",
	})

	if err != nil {
		return nil, err
	}

	var res []*API

	for _, cm := range cms.Items {
		data := cm.Data["api.json"]

		api := &API{}
		err := json.Unmarshal([]byte(data), api)
		if err != nil {
			log.Printf("Unable to parse %s: Error %s", cm.GetObjectMeta().GetName(), err)
			continue
		}
		res = append(res, api)
	}

	return res, nil

}

func getAPIConfig(namespace, apiname string) (*API, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
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
	return api, nil

}

func thisNamespace() (string, error) {

	nsBytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	if err != nil {
		return "", err
	}
	return string(nsBytes), nil
}
