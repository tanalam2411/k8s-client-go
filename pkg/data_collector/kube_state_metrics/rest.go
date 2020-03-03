package kube_state_metrics

import (
	"net/http"
	"net/url"
	"strings"
)

type RESTClient struct {
	base *url.URL
	versionedAPIPath string
	content ClientContentConfig
	Client *http.Client
}


type ClientContentConfig struct {
	AcceptContentTypes string
	ContentType string
}


func NewRESTClient(baseURL *url.URL, versionedAPIPath string, config ClientContentConfig, client *http.Client) (*RESTClient, error) {
	if len(config.ContentType) == 0 {
		config.ContentType = "application/json"
	}

	base := *baseURL
	if !strings.HasSuffix(strings.TrimSpace(base.Path), "/") {
		base.Path =  strings.TrimSpace(base.Path) + "/"
	}

	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		base:             &base,
		versionedAPIPath: versionedAPIPath,
		content:          config,
		Client: client,
	}, nil
}