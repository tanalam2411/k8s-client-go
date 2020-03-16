package main

import (
	"encoding/json"
	"fmt"
	"github.com/tanalam2411/k8s-client-go/pkg/connection"
	"github.com/tanalam2411/k8s-client-go/pkg/data_collector"
)

func main() {

	clientset, err := connection.GetClientSet("/home/afour/.kube/config")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--------- NamesSpaces details: ----------------- ")
	nsds := data_collector.GetNameSpaceDetails(clientset)
	nsdsData, err := nsds.Serialize()
	if err != nil {
		fmt.Println("Failed to serialize NameSpaceDetails: ", err)
	}
	fmt.Println(nsdsData)
	fmt.Println("--------------------------------------------------")

	fmt.Println("--------- Pods per NamesSpace details: -----------")
	ppnsd := data_collector.AllPodsPerNamespace(clientset, nsds.NamespaceNames)
	b, err := json.Marshal(ppnsd)
	if err != nil {
		fmt.Println("Failed to serialize PodsPerNamespaceDetails: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")

	fmt.Println("--------- Deployment per NamesSpace details: -----------")
	dpnsd := data_collector.AllDeploymentsPerNamespace(clientset, nsds.NamespaceNames)
	b, err = json.Marshal(dpnsd)
	if err != nil {
		fmt.Println("Failed to serialize DeploymentsPerNamespaceDetails: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")

	fmt.Println("--------- ReplicaSet per NamesSpace details: -----------")
	rspnsd := data_collector.AllReplicaSetsPerNamespace(clientset, nsds.NamespaceNames)
	b, err = json.Marshal(rspnsd)
	if err != nil {
		fmt.Println("Failed to serialize ReplicaSetPerNamespaceDetails: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")

	fmt.Println("--------- Nodes details: -----------")
	nds := data_collector.GetAllNodesDetails(clientset)
	b, err = json.Marshal(nds)
	if err != nil {
		fmt.Println("Failed to serialize Node Details: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")

	fmt.Println("--------- Service details: -----------")
	spnsd := data_collector.AllServicesPerNamespace(clientset, nsds.NamespaceNames)
	b, err = json.Marshal(spnsd)
	if err != nil {
		fmt.Println("Failed to serialize ServicePerNamespaceDetails: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")

	fmt.Println("--------- ConfigMap details: -----------")
	cmnsd := data_collector.AllConfigMapsPerNamespace(clientset, nsds.NamespaceNames)
	b, err = json.Marshal(cmnsd)
	if err != nil {
		fmt.Println("Failed to serialize ConfigMapsPerNamespaceDetails: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")

	fmt.Println("--------- EndPoints details: -----------")
	epnsd := data_collector.AllEndPointsPerNamespace(clientset, nsds.NamespaceNames)
	b, err = json.Marshal(epnsd)
	if err != nil {
		fmt.Println("Failed to serialize EndPointsPerNamespaceDetails: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")

	fmt.Println("--------- Roles details: -----------")
	rpnsd := data_collector.AllRolesPerNamespace(clientset, nsds.NamespaceNames)
	b, err = json.Marshal(rpnsd)
	if err != nil {
		fmt.Println("Failed to serialize RolesPerNamespaceDetails: ", err)
	}
	fmt.Println(string(b))
	fmt.Println("------------------------------------------------------")
}
