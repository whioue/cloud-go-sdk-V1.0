package v1

import (
	"context"
	"github.com/whioue/cloud-go-sdk-V1.0/domain/v1/tensorboard/dto"
	"github.com/whioue/cloud-go-sdk-V1.0/pkg/json"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type TensorboardGetter interface {
	Tensorboard() ITensorboard
}

type ITensorboard interface {
	Start(c context.Context, args dto.TensorboardStartArgs) (*dto.TensorboardStartReply, error)
	Restart(c context.Context, args dto.TensorboardRestartArgs) error
	Stop(c context.Context, args dto.TensorboardStopArgs) error
	Detail(c context.Context, args dto.TensorboardDetailArgs) (*dto.TensorboardDetailReply, error)
	List(c context.Context, args dto.TensorboardListArgs) (*dto.TensorboardListReply, error)
}

type tensorboard struct {
	client rest.Interface
}

func newTensorboard(c *APIV1Client) *tensorboard {
	return &tensorboard{
		client: c.RESTClient(),
	}
}

func (t *tensorboard) Start(c context.Context, args dto.TensorboardStartArgs) (result *dto.TensorboardStartReply, err error) {
	data, err := t.client.
		Post().
		Resource("task").
		SubResource("tensorboard", "start").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

func (t *tensorboard) Restart(c context.Context, args dto.TensorboardRestartArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("tensorboard", "restart").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t *tensorboard) Stop(c context.Context, args dto.TensorboardStopArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("tensorboard", "stop").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t *tensorboard) Detail(c context.Context, args dto.TensorboardDetailArgs) (result *dto.TensorboardDetailReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("tensorboard").
		VersionedParams(args).
		Do(c).
		Data()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

func (t *tensorboard) List(c context.Context, args dto.TensorboardListArgs) (result *dto.TensorboardListReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("tensorboard", "list").
		VersionedParams(args).
		Do(c).
		Data()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}
