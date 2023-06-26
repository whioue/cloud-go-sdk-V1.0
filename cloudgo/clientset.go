package cloudgo

import (
	"github.com/whioue/cloud-go-sdk-V1.0/cloudgo/service"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type IProject interface {
	ICloud() service.ICloud
}

type ClientSet struct {
	cloud *service.CloudClient
}

func (c *ClientSet) ICloud() service.ICloud {
	return c.cloud
}

func NewForConfig(c *rest.Config) (*ClientSet, error) {
	configShallowCopy := *c

	var (
		cs  ClientSet
		err error
	)

	cs.cloud, err = service.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}

func NewForConfigOrDie(c *rest.Config) *ClientSet {
	var cs ClientSet
	cs.cloud = service.NewForConfigOrDie(c)
	// cs.tms = tms.NewForConfigOrDie(c)
	return &cs
}

func New(c rest.Interface) *ClientSet {
	var cs ClientSet
	cs.cloud = service.New(c)
	// cs.tms = tms.New(c)
	return &cs
}
