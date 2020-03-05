package kbzk8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodInfo contains information about a pod.
type PodInfo struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Labels     map[string]string `json:"labels"`
	Annotaions map[string]string `json:"annotations"`
	Status     v1.PodStatus      `json:"status"`
}

// GetPodInfoList reurns a list of all pods in the cluster.
func GetPodInfoList() ([]*PodInfo, error) {
	cs, err := GetClientSet()

	if err != nil {
		return nil, err
	}
	pl, err := cs.CoreV1().Pods("").List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	var podInfos []*PodInfo

	for _, pod := range pl.Items {

		pi := &PodInfo{
			Name:       pod.GetName(),
			Namespace:  pod.GetNamespace(),
			Labels:     pod.GetLabels(),
			Annotaions: pod.GetAnnotations(),
			Status:     pod.Status,
		}

		podInfos = append(podInfos, pi)

	}
	return podInfos, nil

}
