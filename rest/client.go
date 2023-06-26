package rest

import (
	"net/url"
	"strings"

	"github.com/whioue/cloud-go-sdk-V1.0/pkg/runtime"
	"github.com/whioue/cloud-go-sdk-V1.0/third_party/forked/gorequest"
)

// Interface captures the set of operations for generically interacting with IAM REST apis.
type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Get() *Request
	Delete() *Request
	APIVersion() string
}

// ClientContentConfig controls how RESTClient communicates with the server.
type ClientContentConfig struct {
	BearerToken        string
	AcceptContentTypes string

	ContentType  string
	GroupVersion string
	Negotiator   runtime.ClientNegotiator
}

type RESTClient struct {
	// base is the root URL for all invocations of the client
	base *url.URL
	// group stand for the client group, eg: iam.api, iam.authz
	group string
	// versionedAPIPath is a path segment connecting the base URL to the resource root
	versionedAPIPath string
	// content describes how a RESTClient encodes and decodes responses.
	content ClientContentConfig
	Client  *gorequest.SuperAgent
}

func NewRESTClient(baseURL *url.URL, versionedAPIPath string,
	config ClientContentConfig, client *gorequest.SuperAgent) (*RESTClient, error) {
	if len(config.ContentType) == 0 {
		config.ContentType = "application/json"
	}

	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}

	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		base:             &base,
		group:            config.GroupVersion,
		versionedAPIPath: versionedAPIPath,
		content:          config,
		Client:           client,
	}, nil
}

// Verb begins a Verb request.
func (c *RESTClient) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

// Post begins a POST request. Short for c.Verb("POST").
func (c *RESTClient) Post() *Request {
	return c.Verb("POST")
}

// Put begins a PUT request. Short for c.Verb("PUT").
func (c *RESTClient) Put() *Request {
	return c.Verb("PUT")
}

// Get begins a GET request. Short for c.Verb("GET").
func (c *RESTClient) Get() *Request {
	return c.Verb("GET")
}

// Delete begins a DELETE request. Short for c.Verb("DELETE").
func (c *RESTClient) Delete() *Request {
	return c.Verb("DELETE")
}

// APIVersion returns the APIVersion this RESTClient is expected to use.
func (c *RESTClient) APIVersion() string {
	return c.content.GroupVersion
}
