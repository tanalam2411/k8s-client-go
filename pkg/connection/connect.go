package connection

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var kubeconfig *string

func init() {
	kubeconfig = flag.String("kubeconfig", "", "kubeconfig file")
}

func GetRestclientConfig(configPath string) (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		flag.Set("kubeconfig", configPath)
		flag.Parse()

		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		return config, err
	}

	return config, err
}

func GetClientSet(configPath string) (*kubernetes.Clientset, error) {

	config, err := GetRestclientConfig(configPath)
	if err != nil {
		fmt.Println("Failed to load kubeconfig: #{err}\n")
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Printf("Failed to generate clientset: %f\n", err)
	}
	return clientset, err
}
