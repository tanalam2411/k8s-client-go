package data_collector

import (
	"bytes"
	"fmt"
	"github.com/tanalam2411/k8s-client-go/pkg/connection"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

type Container struct {
	Name string
}

func GetAllContainersOfPod(clientset *kubernetes.Clientset, namespace string, podName string) []*Container {
	pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})

	if err != nil {
		fmt.Println(err)
	}

	var containerList []*Container

	for _, container := range pod.Spec.Containers {
		containerList = append(containerList, &Container{Name: container.Name})
	}

	return containerList
}

func ExecPodContainerCommand(clientset *kubernetes.Clientset, ns string, resourceName string, containerName string) {

	request := clientset.CoreV1().RESTClient().Post().Namespace(ns).Resource("pods").Name(resourceName).SubResource("exec").VersionedParams(&v1.PodExecOptions{
		Stdin:     false,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
		Container: containerName,
		Command:   []string{"ps", "aux"},
	}, scheme.ParameterCodec)

	restConfig, err := connection.GetRestclientConfig("/home/afour/.kube/config")

	if err != nil {
		fmt.Println("Failed to get rest config: ", err)
	}

	executor, err := remotecommand.NewSPDYExecutor(restConfig, "POST", request.URL())
	if err != nil {
		fmt.Println("Failed to get executor: ", err)
	}
	buf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	err = executor.Stream(remotecommand.StreamOptions{
		Stdout: buf,
		Stderr: errBuf,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buf.String())
	fmt.Println(errBuf.String())
}
