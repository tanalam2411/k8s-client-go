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

}
