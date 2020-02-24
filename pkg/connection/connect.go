package connection

import (
	"flag"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClientSet(configPath string) (*kubernetes.Clientset, error) {

	kubeconfig := flag.String("kubeconfig", configPath, "kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		os.Exit(0)
	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Printf("Failed to generate clientset: %f\n", err)
	}

	return clientset, err
}

