package data_collector

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type NameSpaceClient struct {
	Clientset *kubernetes.Clientset
}

type NameSpaceInfo struct {
	Name   string
	UID    types.UID
	Status v1.NamespacePhase
}

type NameSpaceDetails struct {
	NameSpaceInfos []NameSpaceInfo
}

func (nsd *NameSpaceDetails) GetNamespaceDetails(nsl *v1.NamespaceList) {

	for _, item := range nsl.Items {
		nsd.NameSpaceInfos = append(nsd.NameSpaceInfos, NameSpaceInfo{item.Name, item.UID, item.Status.Phase})
	}
}

func (nsc *NameSpaceClient) GetAllNameSpaces() *v1.NamespaceList {
	return nsc.getAllNameSpaces()
}

func (nsc *NameSpaceClient) getAllNameSpaces() *v1.NamespaceList {

	nsl, err := nsc.Clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	return nsl
}
