package data_collector

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type ReplicaSetClient struct {
	Clientset *kubernetes.Clientset
}

type ReplicaSetPerNamespaceDetails struct {
	TotalReplicaSets int32
	Namespace        string
	ReplicaSetsInfo  []ReplicaSetInfo
}

type ReplicaSetInfo struct {
	Name              string
	UID               types.UID
	Namespace         string
	APIVersion        string
	CLusterName       string
	CreationTimestamp metav1.Time
	Replicas          *int32
	ReadyReplicas     int32
	AvailableReplicas int32
}

func GetAllReplicaSetByNamespace(rsc *ReplicaSetClient, namespace string) *ReplicaSetPerNamespaceDetails {

	replicasetList, err := rsc.Clientset.AppsV1().ReplicaSets(namespace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch replicasetList for namespace: ", namespace)
	}

	rspnsd := ReplicaSetPerNamespaceDetails{}
	rspnsd.Namespace = namespace
	rspnsd.TotalReplicaSets = 0

	for _, rs := range replicasetList.Items {
		replicasetInfo := ReplicaSetInfo{
			Name:              rs.GetName(),
			UID:               rs.GetUID(),
			Namespace:         rs.GetNamespace(),
			APIVersion:        rs.APIVersion,
			CLusterName:       rs.GetClusterName(),
			CreationTimestamp: rs.GetCreationTimestamp(),
			Replicas:          rs.Spec.Replicas,
			ReadyReplicas:     rs.Status.ReadyReplicas,
			AvailableReplicas: rs.Status.AvailableReplicas,
		}
		rspnsd.ReplicaSetsInfo = append(rspnsd.ReplicaSetsInfo, replicasetInfo)
		rspnsd.TotalReplicaSets += 1
	}

	return &rspnsd
}

func AllReplicaSetsPerNamespace(clientset *kubernetes.Clientset, namespaces []string) map[string]*ReplicaSetPerNamespaceDetails {
	rspnsd := map[string]*ReplicaSetPerNamespaceDetails{}

	for _, ns := range namespaces {
		rspnsd[ns] = GetAllReplicaSetByNamespace(&ReplicaSetClient{Clientset: clientset}, ns)
	}

	return rspnsd
}
