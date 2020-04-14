package main

import (
	"fmt"
	"github.com/tanalam2411/k8s-client-go/pkg/connection"
	"github.com/tanalam2411/k8s-client-go/pkg/data_collector"
)

func main() {

	clientset, err := connection.GetClientSet("~/.kube/config")
	if err != nil {
		fmt.Println(err)
	}

	nsDetails := data_collector.GetNameSpaceDetails(clientset)

	podsPerNameSpace := data_collector.AllPodsPerNamespace(clientset, nsDetails.NamespaceNames)

	for ns, podList := range podsPerNameSpace {
		fmt.Printf("\n\n\n\n\n --- Namespace: %v-------\n", ns)
		for _, pod := range podList.PodsInfo {
			fmt.Printf("--- Pod: %v\n", pod.Name)
			containers := data_collector.GetAllContainersOfPod(clientset, ns, pod.Name)

			for _, container := range containers {
				fmt.Printf("--- Container: %v\n", container.Name)
				data_collector.ExecPodContainerCommand(clientset, ns, pod.Name, container.Name)
			}
			//data_collector.ExecPodContainerCommand(clientset, ns, pod.Name)
		}
	}

	//data_collector.ExecPodContainerCommand(clientset, "default", "hello-world-f9b447754-fwj5z")
	//data_collector.T1(clientset, "default", "hello-world-f9b447754-fwj5z")
}
