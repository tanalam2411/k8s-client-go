package main

import (
	"fmt"
	"github.com/tanalam2411/k8s-client-go/pkg/connection"
	"github.com/tanalam2411/k8s-client-go/pkg/data_collector"
)

func main() {

	clientset, err := connection.GetClientSet("/home/afour/.kube/config")
	if err != nil {
		fmt.Println(err)
	}

	dc := data_collector.NameSpaceClient{Clientset: clientset}
	nsl := dc.GetAllNameSpaces()

	nsd := data_collector.NameSpaceDetails{}
	nsd.GetNamespaceDetails(nsl)

	// --------- Printing namespaces details --------------------
	fmt.Println("--------- NamesSpaces details: -----------")
	fmt.Println("Name\t\tStatus")
	for _, nsd := range nsd.NameSpaceInfos {
		fmt.Printf("%v\t\t%v\n", nsd.Name, nsd.Status)
	}
	fmt.Println("-------------------------------------------")


	// ----------------------- POD Details ----------------------
	fmt.Println("--------- Pods per NamesSpace details: -----------")
	psc := data_collector.PodClient{Clientset: clientset}
	ppnsd := psc.GetAllPodsByNamespace("kube-system")

	fmt.Println("Namespace: ", ppnsd.Namespace)
	fmt.Println("TotalPods: ", ppnsd.TotalPods)

	for _, pd := range ppnsd.PodsDetails {
		fmt.Println(pd.Name)
		fmt.Println(pd.ClusterName)
		fmt.Println(pd.APIVersion)
		fmt.Println(pd.UID)
	}

	fmt.Println("------------------------------------------------------")
}
