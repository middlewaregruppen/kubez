package api

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//getClientSet returns a client set as a In-Cluster configuration.
func getClientSet() (*kubernetes.Clientset, error) {

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil

}

func getNamespaces() ([]string, error) {

	cs, err := getClientSet()
	if err != nil {
		return nil, err
	}

	nss, err := cs.CoreV1().Namespaces().List(metav1.ListOptions{})

	var res []string
	for _, ns := range nss.Items {
		res = append(res, ns.GetName())
	}

	log.Printf("Namespaces %+v", res)
	return res, nil

}

// getThisImageName returns the image name of the running container.
// If it can not be determined it will return the image docker.io/middlewaregruppen/kubez:latest
func getThisImageName() string {

	defaultImage := "docker.io/middlewaregruppen/kubez:latest"

	cs, err := getClientSet()
	if err != nil {
		return defaultImage
	}

	// Get the hostname of this deployment.
	podname, err := os.Hostname()

	if err != nil {
		return defaultImage
	}

	// Namespace
	ns, err := thisNamespace()
	if err != nil {
		return defaultImage
	}

	// Fetch pod
	pod, err := cs.CoreV1().Pods(ns).Get(podname, metav1.GetOptions{})
	if err != nil {
		return defaultImage
	}

	// One image configured?
	if len(pod.Spec.Containers) == 1 {
		return pod.Spec.Containers[0].Image
	}

	//Multiple images configured
	if len(pod.Spec.Containers) > 1 {
		for _, c := range pod.Spec.Containers {

			if strings.Contains(c.Image, "kubez") {
				return c.Image
			}
		}
	}

	return defaultImage
}

func thisNamespace() (string, error) {

	nsBytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	if err != nil {
		return "", err
	}
	return string(nsBytes), nil
}
