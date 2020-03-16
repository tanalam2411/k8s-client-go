package data_collector

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type EndPointsClient struct {
	Clientset *kubernetes.Clientset
}

type EndPointsPerNamespaceDetails struct {
	Namespace     string
	EndpointsInfo []EndPointsInfo
}

type EndPointsInfo struct {
	Name              string
	UID               types.UID
	Namespace         string
	APIVersion        string
	ClusterName       string
	CreationTimestamp metav1.Time
	Subsets           []v1.EndpointSubset
}

func GetAllEndPointsByNamespace(epc *EndPointsClient, namespace string) *EndPointsPerNamespaceDetails {

	endpointList, err := epc.Clientset.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch endpointsList for namespace: ", namespace)
	}

	eppnsd := EndPointsPerNamespaceDetails{}
	eppnsd.Namespace = namespace

	for _, ep := range endpointList.Items {
		endpointInfo := EndPointsInfo{
			Name:              ep.GetName(),
			UID:               ep.GetUID(),
			Namespace:         ep.GetNamespace(),
			APIVersion:        ep.APIVersion,
			ClusterName:       ep.GetClusterName(),
			CreationTimestamp: ep.GetCreationTimestamp(),
			Subsets:           ep.Subsets,
		}
		eppnsd.EndpointsInfo = append(eppnsd.EndpointsInfo, endpointInfo)
	}

	return &eppnsd
}

func AllEndPointsPerNamespace(clientset *kubernetes.Clientset, namespace []string) map[string]*EndPointsPerNamespaceDetails {
	eppnsd := map[string]*EndPointsPerNamespaceDetails{}

	for _, ns := range namespace {
		eppnsd[ns] = GetAllEndPointsByNamespace(&EndPointsClient{Clientset: clientset}, ns)
	}

	return eppnsd
}
