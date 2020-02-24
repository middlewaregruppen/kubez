package kbzk8s

import (
	"fmt"
	"io/ioutil"

	petname "github.com/dustinkirkland/golang-petname"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// K8SLoad is the configuration for Load().
//    Namespace 		The namespace to create the load in. If empty the load will be created in the namespace the pod runs in.
//    Deployments		Number of deployments
//    Pods				Number of pods to create in the namespace.
//    RequestedCPU  	Set resources in the pod specification.
//						eg. 300m
//						If 0 or empty it will be omitted in the deployment spec.
//    RequestedMemory 	Set resources in the pod specification.
//						eg. 200Mi
//						If 0 or empty it will be omitted in the deployment spec.
//    LimitCPU			Set resources in the pod specification.
//						eg. 300m
//						If 0 or empty it will be omitted in the deployment spec.
//    LimitMemory		Set resources in the pod specification.
//						eg. 300Mi
//						If 0 or empty it will be omitted in the deployment spec.
//    LoadProfile One of the following:
//       - "none", does nothing
//       - "cpu", will load the cpu 100%
//       - "mem200", will allocate 200 Mb Ram
//       - "mem2000", will allocate 2000 Mb Ram.
//		 - "log100ms", write a log line and waits 100ms (~10 lines/s)
//       - "log9ms" write a log line and waits 9ms (~100 lines/s)
//       - "log3ms" write a log line and waits 3ms (~300 lines/s)

type K8SLoad struct {
	Namespace       string `json:"namespace"`
	Deployments     int    `json:"deployments,string"`
	Pods            int32  `json:"pods,string"`
	RequestedCPU    string `json:"reqCPU"`
	RequestedMemory string `json:"reqMem"`
	LimitCPU        string `json:"limCPU"`
	LimitMemory     string `json:"limMem"`
	LoadProfile     string `json:"profile"`
}

// Load creates deployments according to K8SLoad
func Load(load *K8SLoad) error {

	if load.Namespace == "" {
		nsBytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
		if err != nil {
			return err
		}
		load.Namespace = string(nsBytes)

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

	// Limits and requests in podspec.
	limits := make(v1.ResourceList)
	if len(load.LimitCPU) > 0 && load.LimitCPU != "0" {
		limits["cpu"], _ = resource.ParseQuantity(load.LimitCPU)
	}
	if len(load.LimitMemory) > 0 && load.LimitMemory != "0" {
		limits["memory"], _ = resource.ParseQuantity(load.LimitMemory)
	}

	requests := make(v1.ResourceList)
	if len(load.RequestedCPU) > 0 && load.RequestedCPU != "0" {
		requests["cpu"], _ = resource.ParseQuantity(load.RequestedCPU)
	}
	if len(load.RequestedMemory) > 0 && load.RequestedMemory != "0" {
		requests["memory"], _ = resource.ParseQuantity(load.RequestedMemory)
	}

	// Load Profile
	command := []string{}
	args := []string{}

	switch load.LoadProfile {
	case "none":
		command = []string{"/loader"}
		args = []string{"-profile", "none"}

	case "cpu":
		command = []string{"/loader"}
		args = []string{"-profile", "cpu"}

	case "mem100":
		command = []string{"/loader"}
		args = []string{"-profile", "mem", "-mb", "100"}

	case "mem2000":
		command = []string{"/loader"}
		args = []string{"-profile", "mem", "-mb", "2000"}

	case "log100ms":
		command = []string{"/loader"}
		args = []string{"-profile", "log", "-logwait", "100ms"}

	case "log9ms":
		command = []string{"/loader"}
		args = []string{"-profile", "log", "-logwait", "9ms"}

	case "log3ms":
		command = []string{"/loader"}
		args = []string{"-profile", "lgo", "-logwait", "3ms"}
	}

	for d := 0; d < load.Deployments; d++ {

		dpNmae := fmt.Sprintf("kl-%s", petname.Generate(3, "-"))

		clientset.AppsV1().Deployments(load.Namespace).Create(&appv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: dpNmae,
			},
			Spec: appv1.DeploymentSpec{
				Replicas: &load.Pods,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"app": dpNmae},
				},
				Template: v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": dpNmae}},
					Spec: v1.PodSpec{
						Containers: []v1.Container{

							{
								Resources: v1.ResourceRequirements{
									Limits:   limits,
									Requests: requests,
								},
								Name:  "kubez",
								Image: "docker.io/middlewaregruppen/kubez",
								Ports: []v1.ContainerPort{
									{ContainerPort: 3000},
								},
								Command: command,
								Args:    args,
							},
						},
					},
				},
			},
		})

	}
	return nil
}
