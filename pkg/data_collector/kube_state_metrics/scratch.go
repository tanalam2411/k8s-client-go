package kube_state_metrics

import (
	"fmt"
	"io"
	"os"

	"encoding/json"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/prom2json"
	"log"
	"net/http"
)

// attempt to get metrics from kube state metrics and parse using prom to json lib

var metricsName = map[string]bool{"kube_pod_container_resource_requests": true,
	"kube_pod_container_resource_requests_memory_bytes": true}

func GetMetrics() {

	requestClient := http.Client{}

	resp, err := requestClient.Get("http://127.0.0.1:8080/metrics")

	if err != nil {
		fmt.Println("Failed to get metrics: ", err)
	}

	Prom2Json(resp.Body)
}

func Prom2Json(input io.Reader) {

	mfChan := make(chan *dto.MetricFamily, 1024)

	if err := prom2json.ParseReader(input, mfChan); err != nil {
		log.Fatal("error reading metrics:", err)
	}

	result := map[string][]*prom2json.Family{}
	for mf := range mfChan {

		if metricsName[*mf.Name] {
			fmt.Println(*mf.Name)
			result[*mf.Name] = append(result[*mf.Name], prom2json.NewFamily(mf))
		}
	}

	fmt.Println(len(result))
	jsonText, err := json.Marshal(result)
	if err != nil {
		log.Fatalln("error marshaling JSON:", err)
	}
	if _, err := os.Stdout.Write(jsonText); err != nil {
		log.Fatalln("error writing to stdout:", err)
	}
	fmt.Println()
}
