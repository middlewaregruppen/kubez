package kbzk8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodInfo contains information about a pod.
type PodInfo struct {
	Name              string                    `json:"name"`
	Namespace         string                    `json:"namespace"`
	Labels            map[string]string         `json:"labels"`
	Annotaions        map[string]string         `json:"annotations"`
	Status            v1.PodStatus              `json:"status"`
	Condition         string                    `json:"condition"`
	ContainerRestarts int32                     `json:"containerRestarts"`
	ContainerInfo     map[string]*ContainerInfo `json:"containerInfo"`
}

// ContainerInfo contains a subset of information about the containers in the Pod.
type ContainerInfo struct {
	LivenessProbeConfig        *v1.Probe                `json:"livenessProbeConfig"`
	RedynessProbeConfig        *v1.Probe                `json:"redynessProbeConfig"`
	ResourceRequirementsConfig *v1.ResourceRequirements `json:"resouceRequirementsConfig"`
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
			Name:              pod.GetName(),
			Namespace:         pod.GetNamespace(),
			Labels:            pod.GetLabels(),
			Annotaions:        pod.GetAnnotations(),
			Status:            pod.Status,
			Condition:         condition(pod.Status),
			ContainerRestarts: restarts(pod.Status),
			ContainerInfo:     make(map[string]*ContainerInfo),
		}

		for _, c := range pod.Spec.Containers {

			ci := &ContainerInfo{
				LivenessProbeConfig:        c.LivenessProbe,
				RedynessProbeConfig:        c.ReadinessProbe,
				ResourceRequirementsConfig: &c.Resources,
			}

			pi.ContainerInfo[c.Name] = ci
		}

		podInfos = append(podInfos, pi)

	}
	return podInfos, nil

}

// condition returns the Ready Condition
func condition(status v1.PodStatus) string {
	for _, c := range status.Conditions {
		if c.Type == v1.PodReady {
			switch c.Status {
			case v1.ConditionFalse:
				return c.Reason
			case v1.ConditionUnknown:
				return "Unknown"
			case v1.ConditionTrue:
				return "Ready"
			}
		}
	}
	return "Unknown"
}

// summorizes the restarts
func restarts(status v1.PodStatus) int32 {
	var restarts int32
	for _, c := range status.ContainerStatuses {
		restarts = restarts + c.RestartCount
	}
	return restarts
}
