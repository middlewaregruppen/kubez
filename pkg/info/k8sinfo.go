package info

import (
	"io/ioutil"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8SInfo struct {
	Namespace                      string `json:"namespace"`
	NumberOfNamespaces             int    `json:"noNamespaces"`
	NumberOfPods                   int    `json:"noPods"`
	NumberOfDeployments            int    `json:"noDeployments"`
	NumberOfReadyPods              int    `json:"noReadyPods"`
	NumberOfPodsInNamespace        int    `json:"noPodsInNs"`
	NumberOfReadyPodsInNamespace   int    `json:"noReadyPodsInNs"`
	NumberOfDeploymentsInNamespace int    `json:"noDeploymentsinNs"`
}

func GetK8SInfo() *K8SInfo {
	ki := &K8SInfo{}

	nsb, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	ki.Namespace = string(nsb)

	if err != nil {
		log.Printf("%s", err)
		return ki
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("%s", err)
		return ki
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("%s", err)
		return ki
	}

	// NAMESPACES
	nslist, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {
		ki.NumberOfNamespaces = len(nslist.Items)
	}

	// POD in Cluster
	podlist, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {
		ki.NumberOfPods = len(podlist.Items)
		// POD READY in Cluster
		for _, pod := range podlist.Items {
			if len(pod.Status.ContainerStatuses) > 0 {
				if pod.Status.ContainerStatuses[0].Ready {
					ki.NumberOfReadyPods++
				}
			}
		}
	}
	// POD in namespace
	podlist, err = clientset.CoreV1().Pods(ki.Namespace).List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {
		ki.NumberOfPodsInNamespace = len(podlist.Items)
		// POD READY in Cluster
		for _, pod := range podlist.Items {
			if pod.Status.ContainerStatuses[0].Ready {
				ki.NumberOfReadyPodsInNamespace++
			}
		}
	}

	// Deployments in Cluster

	dlist, err := clientset.AppsV1().Deployments("").List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {

		ki.NumberOfDeployments = len(dlist.Items)
	}

	// Deployments in NS

	dlist, err = clientset.AppsV1().Deployments(ki.Namespace).List(metav1.ListOptions{})

	// Ignore errors.
	if err == nil {

		ki.NumberOfDeploymentsInNamespace = len(dlist.Items)
	}

	return ki

}
