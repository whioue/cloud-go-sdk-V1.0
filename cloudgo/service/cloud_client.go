package service

import (
	v1 "github.com/whioue/cloud-go-sdk-V1.0/cloudgo/service/v1"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type ICloud interface {
	ApiV1() v1.IApiV1
}

type CloudClient struct {
	apiV1 *v1.APIV1Client
}

func (c *CloudClient) ApiV1() v1.IApiV1 {
	return c.apiV1
}

func NewForConfig(c *rest.Config) (*CloudClient, error) {
	configShallowCopy := *c

	var ic CloudClient

	var err error

	ic.apiV1, err = v1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &ic, nil
}

func NewForConfigOrDie(c *rest.Config) *CloudClient {
	var ic CloudClient
	ic.apiV1 = v1.NewForConfigOrDie(c)
	return &ic
}

func New(c rest.Interface) *CloudClient {
	var ic CloudClient
	ic.apiV1 = v1.New(c)
	return &ic
}
