package data_collector

import (
	"fmt"
	apiV1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type ServiceClient struct {
	Clientset *kubernetes.Clientset
}

type ServicePerNamespaceDetails struct {
	Namespace    string
	ServicesInfo []ServiceInfo
}

type ServiceInfo struct {
	Name              string
	UID               types.UID
	Namespace         string
	APIVersion        string
	ClusterName       string
	CreationTimestamp metav1.Time
	Type              string
	Ports             []apiV1.ServicePort
	ClusterIP         string
	ExternalIPs       []string
}

func GetAllServiceByNamespace(sc *ServiceClient, namespace string) *ServicePerNamespaceDetails {

	serviceList, err := sc.Clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch serviceList for namespace: ", namespace)
	}

	spnsd := ServicePerNamespaceDetails{}
	spnsd.Namespace = namespace

	for _, s := range serviceList.Items {
		serviceInfo := ServiceInfo{
			Name:              s.GetName(),
			UID:               s.GetUID(),
			Namespace:         s.GetNamespace(),
			APIVersion:        s.APIVersion,
			ClusterName:       s.GetClusterName(),
			CreationTimestamp: s.GetCreationTimestamp(),
			Type:              string(s.Spec.Type),
			Ports:             s.Spec.Ports,
			ClusterIP:         s.Spec.ClusterIP,
			ExternalIPs:       s.Spec.ExternalIPs,
		}
		spnsd.ServicesInfo = append(spnsd.ServicesInfo, serviceInfo)
	}

	return &spnsd
}

func AllServicesPerNamespace(clietset *kubernetes.Clientset, namespaces []string) map[string]*ServicePerNamespaceDetails {
	spnsd := map[string]*ServicePerNamespaceDetails{}

	for _, ns := range namespaces {
		spnsd[ns] = GetAllServiceByNamespace(&ServiceClient{Clientset: clietset}, ns)
	}

	return spnsd
}
