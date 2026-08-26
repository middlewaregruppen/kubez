package api

import (
	"log"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// getClientSet builds a Kubernetes clientset. When running inside a pod it
// uses the in-cluster service-account credentials. When running locally it
// falls back to the kubeconfig file (~/.kube/config or $KUBECONFIG).
func getClientSet() (*kubernetes.Clientset, error) {

	config, err := rest.InClusterConfig()
	if err != nil {
		// Not running in a pod — fall back to local kubeconfig.
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		config, err = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			loadingRules,
			&clientcmd.ConfigOverrides{},
		).ClientConfig()
		if err != nil {
			return nil, err
		}
	}

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
	if err != nil {
		return nil, err
	}

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

	// In-cluster: namespace is injected as a file by the kubelet.
	nsBytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err == nil {
		return string(nsBytes), nil
	}

	// Local: derive the namespace from the active kubeconfig context.
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	ns, _, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	).Namespace()
	if err != nil {
		return "", err
	}
	return ns, nil
}
