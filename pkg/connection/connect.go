package connection

import (
	//"encoding/json"
	"flag"
	"fmt"
	"os"

	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

//func Connect() {
//
//	kubeconfig := flag.String("kubeconfig", "/home/afour/.kube/config", "kubecnfig file")
//	flag.Parse()
//
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
//		os.Exit(0)
//	}
//	clientset, err := kubernetes.NewForConfig(config)
//
//	pod, err := clientset.CoreV1().Pods("kube-system").Get("etcd-kind-control-plane", metav1.GetOptions{})
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	pod_data, err := json.Marshal(pod)
//
//	fmt.Println(err)
//	fmt.Println(string(pod_data))
//
//}
