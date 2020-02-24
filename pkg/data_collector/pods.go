package data_collector

import (
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1")


type PodClient struct {
	Clientset *kubernetes.Clientset
}


type PodsPerNameSpaceDetails struct {
	TotalPods	int
	Namespace 	string
	PodsDetails  []PodDetails
}


type PodDetails struct{
	Name 		string
	UID			types.UID
	Namespace	string
	APIVersion 	string
	ClusterName string
}


func (pc *PodClient) GetAllPodsByNamespace(nameSpace string) *PodsPerNameSpaceDetails {

	podList, err := pc.Clientset.CoreV1().Pods(nameSpace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch podList for namespace: ", nameSpace)
		fmt.Println(err)
	}

	ppsd := PodsPerNameSpaceDetails{}
	ppsd.Namespace = nameSpace
	ppsd.TotalPods = 0
	for _, pod := range podList.Items {
		fmt.Println(pod.Name)
		podDetails := PodDetails{Name: pod.Name,
								UID: pod.UID,
								Namespace: pod.Namespace,
								APIVersion: pod.APIVersion,
								ClusterName: pod.ClusterName,
								}
		ppsd.PodsDetails = append(ppsd.PodsDetails, podDetails)
		ppsd.TotalPods += 1
	}

	return &ppsd
}
