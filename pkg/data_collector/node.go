package data_collector

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type NodeClient struct {
	Clientset *kubernetes.Clientset
}

type NodeDetails struct {
	NodeInfo       *NodeInfo
	NodeAddresses  []NodeAddress
	NodeSystemInfo *NodeSystemInfo
	NodeCapacity   *NodeCapacity
}

type NodeInfo struct {
	Name string
	UID  types.UID
}

func (ni *NodeInfo) setter(node *v1.Node) {
	ni.Name = node.GetName()
	ni.UID = node.GetUID()
}

type NodeAddress struct {
	Address string
	Type    string
}

func (na *NodeAddress) setter(nodeAddr *v1.NodeAddress) {
	na.Address = nodeAddr.Address
	na.Type = string(nodeAddr.Type)
}

type NodeSystemInfo struct {
	Architecture     string
	KernelVersion    string
	KubeProxyVersion string
	KubeletVersion   string
	OperatingSystem  string
	OsImage          string
	SystemUUID       string
}

func (nsi *NodeSystemInfo) setter(node *v1.Node) {
	nsi.Architecture = node.Status.NodeInfo.Architecture
	nsi.KernelVersion = node.Status.NodeInfo.KernelVersion
	nsi.OperatingSystem = node.Status.NodeInfo.OperatingSystem
	nsi.OsImage = node.Status.NodeInfo.OSImage
	nsi.SystemUUID = node.Status.NodeInfo.SystemUUID
	nsi.KubeletVersion = node.Status.NodeInfo.KernelVersion
	nsi.KubeProxyVersion = node.Status.NodeInfo.KubeProxyVersion
}

type MaxPodsPerNode int8

type NodeCapacity struct {
	CPU              string
	Memory           string
	EphemeralStorage string
	Pods             MaxPodsPerNode
}

func (nd *NodeCapacity) setter(node *v1.Node) {
	nd.CPU = node.Status.Capacity.Cpu().String()
	nd.Memory = node.Status.Capacity.Memory().String()
	nd.EphemeralStorage = node.Status.Capacity.StorageEphemeral().String()
	nd.Pods = MaxPodsPerNode(node.Status.Capacity.Pods().Size())
}

func getAllNodes(nc *NodeClient) (*v1.NodeList, error) {
	nodeList, err := nc.Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	return nodeList, err
}

func GetAllNodesDetails(clientset *kubernetes.Clientset) map[string]*NodeDetails {
	nodeList, err := getAllNodes(&NodeClient{Clientset: clientset})

	if err != nil {
		fmt.Println("Failed to get node list: ", err)
	}

	nodeMap := map[string]*NodeDetails{}

	for _, node := range nodeList.Items {
		nc := NodeDetails{}
		nc.NodeInfo = &NodeInfo{}
		nc.NodeInfo.setter(&node)
		nc.NodeSystemInfo = &NodeSystemInfo{}
		nc.NodeSystemInfo.setter(&node)
		nc.NodeCapacity = &NodeCapacity{}
		nc.NodeCapacity.setter(&node)
		for _, addr := range node.Status.Addresses {
			na := NodeAddress{Address: addr.Address, Type: string(addr.Type)}
			nc.NodeAddresses = append(nc.NodeAddresses, na)
		}

		nodeMap[node.Name] = &nc
	}

	return nodeMap
}
