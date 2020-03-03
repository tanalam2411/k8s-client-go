package main

import "github.com/tanalam2411/k8s-client-go/pkg/data_collector/kube_state_metrics"

func main() {
	kube_state_metrics.GetMetrics()
}
