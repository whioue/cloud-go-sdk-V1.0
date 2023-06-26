package rest

import (
	"fmt"
	"github.com/whioue/cloud-go-sdk-V1.0/pkg/runtime"
	"net/http"
	"time"

	"github.com/whioue/cloud-go-sdk-V1.0/third_party/forked/gorequest"
)

type Config struct {
	Host    string
	APIPath string
	ContentConfig

	BearerToken string

	// The maximum length of time to wait before giving up on a server request. A value of zero means no timeout.
	Timeout       time.Duration
	MaxRetries    int
	RetryInterval time.Duration
}

// ContentConfig defines config for content.
type ContentConfig struct {
	ServiceName        string
	AcceptContentTypes string
	ContentType        string
	GroupVersion       string
	Negotiator         runtime.ClientNegotiator
}

func RESTClientFor(config *Config) (*RESTClient, error) {
	if config.GroupVersion == "" {
		return nil, fmt.Errorf("GroupVersion is required when initializing a RESTClient")
	}

	baseURL, versionedAPIPath, err := defaultServerURLFor(config)
	if err != nil {
		return nil, err
	}

	// Only retry when get a server side error.
	client := gorequest.New().Timeout(config.Timeout).
		Retry(config.MaxRetries, config.RetryInterval, http.StatusInternalServerError)
	// NOTICE: must set DoNotClearSuperAgent to true, or the client will clean header befor http.Do
	client.DoNotClearSuperAgent = true

	var gv string
	if config.GroupVersion != "" {
		gv = config.GroupVersion
	}

	clientContent := ClientContentConfig{
		BearerToken:        config.BearerToken,
		AcceptContentTypes: config.AcceptContentTypes,
		ContentType:        config.ContentType,
		GroupVersion:       gv,
		Negotiator:         config.Negotiator,
	}

	return NewRESTClient(baseURL, versionedAPIPath, clientContent, client)
}
