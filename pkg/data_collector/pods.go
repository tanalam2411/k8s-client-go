package data_collector

import (
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type PodClient struct {
	Clientset *kubernetes.Clientset
}

type PodsPerNameSpaceDetails struct {
	TotalPods int32
	Namespace string
	PodsInfo  []PodInfo
}

type PodInfo struct {
	Name              string
	UID               types.UID
	Namespace         string
	APIVersion        string
	ClusterName       string
	CreationTimestamp metav1.Time
}

func (ppnsd *PodsPerNameSpaceDetails) Serialize() (string, error) {
	ppnsdBytes, err := json.Marshal(ppnsd)
	return string(ppnsdBytes), err
}

func GetAllPodsByNamespace(pc *PodClient, namespace string) *PodsPerNameSpaceDetails {

	podList, err := pc.Clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch podList for namespace: ", namespace)
		fmt.Println(err)
	}

	ppsd := PodsPerNameSpaceDetails{}
	ppsd.Namespace = namespace
	ppsd.TotalPods = 0
	for _, pod := range podList.Items {
		podInfo := PodInfo{Name: pod.Name,
			UID:               pod.UID,
			Namespace:         pod.Namespace,
			APIVersion:        pod.APIVersion,
			ClusterName:       pod.ClusterName,
			CreationTimestamp: pod.CreationTimestamp,
		}
		ppsd.PodsInfo = append(ppsd.PodsInfo, podInfo)
		ppsd.TotalPods += 1
	}

	return &ppsd
}

func AllPodsPerNamespace(clientset *kubernetes.Clientset, namespaces []string) map[string]*PodsPerNameSpaceDetails {
	ppnsd := map[string]*PodsPerNameSpaceDetails{}
	for _, ns := range namespaces {
		ppnsd[ns] = GetAllPodsByNamespace(&PodClient{clientset}, ns)

	}
	return ppnsd
}
