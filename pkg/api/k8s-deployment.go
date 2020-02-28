package api

import (
	"log"

	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//listAPIEndpointDeployments returns all the APIEndpoints in the cluster.
func getAPIEndpointDeployments() ([]appv1.Deployment, error) {

	clientset, err := getClientSet()
	if err != nil {
		return nil, err
	}
	var res []appv1.Deployment

	nsl, err := getNamespaces()
	if err != nil {
		return nil, err
	}

	for _, ns := range nsl {

		dlst, err := clientset.AppsV1().Deployments(ns).List(metav1.ListOptions{
			LabelSelector: "kubez-type=api-endpoint",
		})

		if err != nil {
			log.Printf("Failure getting deployemnts from: %s %s", ns, err)
			continue
		}

		for _, d := range dlst.Items {
			res = append(res, d)
		}
	}

	return res, nil
}

//createAPIEndpointDeployment creates an new deployment for an api-endpoint.
func createAPIEndpointDeployment(namespace, deploymentName string, api *API) error {

	clientset, err := getClientSet()
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
							Image: getThisImageName(),
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
