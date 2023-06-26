package v1

import (
	"github.com/whioue/cloud-go-sdk-V1.0/pkg/runtime"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type IApiV1 interface {
	RESTClient() rest.Interface
	LogGetter
	ModelGetter
	TaskGetter
	TensorboardGetter
	TrainGetter
}

type APIV1Client struct {
	restClient rest.Interface
}

func (c *APIV1Client) Log() ILog {
	return newLog(c)
}

func (c *APIV1Client) Model() IModel {
	return newModel(c)
}

func (c *APIV1Client) Task() ITask {
	return newTask(c)
}

func (c *APIV1Client) Tensorboard() ITensorboard {
	return newTensorboard(c)
}

func (c *APIV1Client) Train() ITrain {
	return newTrain(c)
}

// NewForConfig creates a new APIV1Client for the given config.
func NewForConfig(c *rest.Config) (*APIV1Client, error) {
	config := *c
	config.Negotiator = runtime.NewSimpleClientNegotiator()
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &APIV1Client{client}, nil
}

// NewForConfigOrDie creates a new APIV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *APIV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}

	return client
}

// New creates a new APIV1Client for the given RESTClient.
func New(c rest.Interface) *APIV1Client {
	return &APIV1Client{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *APIV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}

	return c.restClient
}
