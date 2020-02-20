package info

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8SInfo struct {
	NumberOfNamespaces  int `json:"noNamespaces"`
	NumberOfPods        int `json:"noPods"`
	NumberOfDeployments int `json:"noDeployments"`
	NumberOfReadyPods   int `json:"noReadyPods"`
}

func GetK8SInfo() *K8SInfo {
	ki := &K8SInfo{}

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// NAMESPACES
	nslist, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {

		ki.NumberOfNamespaces = len(nslist.Items)
	}

	// POD
	podlist, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {

		ki.NumberOfPods = len(podlist.Items)
	}

	// POD READY
	for _, pod := range podlist.Items {
		if pod.Status.ContainerStatuses[0].Ready {
			ki.NumberOfReadyPods++
		}
	}

	// Deployments

	dlist, err := clientset.AppsV1().Deployments("").List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {

		ki.NumberOfDeployments = len(dlist.Items)
	}

	return ki

}
