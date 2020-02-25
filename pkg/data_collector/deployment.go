package data_collector

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type DeploymentClient struct {
	Clientset *kubernetes.Clientset
}

type DeploymentPerNamespaceDetails struct {
	TotalDeployments int32
	Namespace string
	DeploymentsInfo []DeploymentInfo
}

type DeploymentInfo struct {
	Name              string
	UID               types.UID
	Namespace         string
	APIVersion        string
	ClusterName       string
	CreationTimestamp metav1.Time
	Replicas          *int32
}


func GetAllDeploymentsByNamespace(dc *DeploymentClient, namespace string) *DeploymentPerNamespaceDetails {

	deploymentList, err := dc.Clientset.AppsV1().Deployments(namespace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch deploymentList for namespace: ", namespace)
		fmt.Println(err)
	}

	dpnsd := DeploymentPerNamespaceDetails{}
	dpnsd.Namespace = namespace
	dpnsd.TotalDeployments = 0

	for _, deployment := range deploymentList.Items {
		deploymentInfo := DeploymentInfo{
			Name:              deployment.Name,
			UID:               deployment.UID,
			Namespace:         deployment.Namespace,
			APIVersion:        deployment.APIVersion,
			ClusterName:       deployment.ClusterName,
			CreationTimestamp: deployment.CreationTimestamp,
			Replicas:          deployment.Spec.Replicas,
		}
		dpnsd.DeploymentsInfo = append(dpnsd.DeploymentsInfo, deploymentInfo)
		dpnsd.TotalDeployments += 1
	}

	return &dpnsd
}


func AllDeploymentsPerNamespace(clientset *kubernetes.Clientset, namespaces []string) map[string]*DeploymentPerNamespaceDetails {
	dpnsd := map[string]*DeploymentPerNamespaceDetails{}

	for _, ns := range namespaces {
		dpnsd[ns] = GetAllDeploymentsByNamespace(&DeploymentClient{clientset}, ns)
	}

	return dpnsd
}
