package v1

import (
	"context"
	"github.com/whioue/cloud-go-sdk-V1.0/domain/v1/model/dto"
	"github.com/whioue/cloud-go-sdk-V1.0/pkg/json"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type ModelGetter interface {
	Model() IModel
}

type IModel interface {
	Delete(c context.Context, args dto.ModelDeleteArgs) error
	Catalogs(c context.Context, args dto.ModelCatalogsArgs) (*dto.ModelCatalogsReply, error)
	Download(c context.Context, args dto.ModelDownloadArgs) (*dto.ModelDownloadReply, error)
	List(c context.Context, args dto.ModelListArgs) (*dto.ModelListReply, error)
}

type model struct {
	client rest.Interface
}

func newModel(c *APIV1Client) *model {
	return &model{
		client: c.RESTClient(),
	}
}

func (m *model) Delete(c context.Context, args dto.ModelDeleteArgs) error {
	return nil
}

func (m *model) Catalogs(c context.Context, args dto.ModelCatalogsArgs) (result *dto.ModelCatalogsReply, err error) {
	data, err := m.client.
		Get().
		Resource("model").
		SubResource("catalogs").
		VersionedParams(args).
		Do(c).
		Data()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

func (m *model) Download(c context.Context, args dto.ModelDownloadArgs) (result *dto.ModelDownloadReply, err error) {
	return nil, nil
}

func (m *model) List(c context.Context, args dto.ModelListArgs) (result *dto.ModelListReply, err error) {
	data, err := m.client.
		Get().
		Resource("model").
		SubResource("list").
		VersionedParams(args).
		Do(c).
		Data()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}
