package main

import (
	"github.com/tanalam2411/k8s-client-go/pkg/data_collector/blackbox_exporter"
	"net/url"
)


func main() {

	_url := &url.URL{}
	blackbox_exporter.HttpProbe(_url, "http://localhost:9115")
}