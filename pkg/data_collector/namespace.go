package data_collector

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type NamespaceClient struct {
	Clientset *kubernetes.Clientset
}

type NamespaceInfo struct {
	Name              string
	UID               types.UID
	Status            v1.NamespacePhase
	CreationTimestamp metav1.Time
}

type NamespaceDetails struct {
	NamespacesInfo []NamespaceInfo
	NamespaceNames []string
}

func (nsd *NamespaceDetails) Serialize() (string, error) {
	bytes, err := json.Marshal(nsd)
	return string(bytes), err
}

func (nsd *NamespaceDetails) GetNamespaceDetails(nsl *v1.NamespaceList) {

	for _, item := range nsl.Items {
		nsd.NamespacesInfo = append(nsd.NamespacesInfo, NamespaceInfo{
			Name:              item.Name,
			UID:               item.UID,
			Status:            item.Status.Phase,
			CreationTimestamp: metav1.Time{},
		})

		nsd.NamespaceNames = append(nsd.NamespaceNames, item.Name)
	}
}

func GetAllNamespaces(nsc *NamespaceClient) *v1.NamespaceList {
	return getAllNamespaces(nsc)
}

func getAllNamespaces(nsc *NamespaceClient) *v1.NamespaceList {

	nsl, err := nsc.Clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	return nsl
}

func GetNameSpaceDetails(clientset *kubernetes.Clientset) *NamespaceDetails {
	nsl := getAllNamespaces(&NamespaceClient{clientset})

	nsd := NamespaceDetails{}
	nsd.GetNamespaceDetails(nsl)

	return &nsd
}
