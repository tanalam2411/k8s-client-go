package data_collector

import (
	"fmt"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type RolesClient struct {
	Clientset *kubernetes.Clientset
}

type RolesPerNamespaceDetails struct {
	Namespace string
	RolesInfo []RolesInfo
}

type RolesInfo struct {
	Name              string
	UID               types.UID
	Namespace         string
	APIVersion        string
	Clustername       string
	CreationTimestamp metav1.Time
	Rules             []v1.PolicyRule
}

func GetAllRolesByNamespace(rc *RolesClient, namespace string) *RolesPerNamespaceDetails {

	rolesList, err := rc.Clientset.RbacV1().Roles(namespace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch rolesList for namespace: ", namespace)
	}

	rpnsd := RolesPerNamespaceDetails{}
	rpnsd.Namespace = namespace

	for _, r := range rolesList.Items {
		rolesInfo := RolesInfo{
			Name:              r.GetName(),
			UID:               r.GetUID(),
			Namespace:         r.GetNamespace(),
			APIVersion:        r.APIVersion,
			Clustername:       r.GetClusterName(),
			CreationTimestamp: r.GetCreationTimestamp(),
			Rules:             r.Rules,
		}
		rpnsd.RolesInfo = append(rpnsd.RolesInfo, rolesInfo)
	}
	return &rpnsd
}

func AllRolesPerNamespace(clientset *kubernetes.Clientset, namespaces []string) map[string]*RolesPerNamespaceDetails {
	rpnsd := map[string]*RolesPerNamespaceDetails{}

	for _, ns := range namespaces {
		rpnsd[ns] = GetAllRolesByNamespace(&RolesClient{Clientset: clientset}, ns)
	}
	return rpnsd
}
