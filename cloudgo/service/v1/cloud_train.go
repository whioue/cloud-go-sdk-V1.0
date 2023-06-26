package v1

import (
	"context"
	"github.com/whioue/cloud-go-sdk-V1.0/domain/v1/train/dto"
	"github.com/whioue/cloud-go-sdk-V1.0/pkg/json"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type TrainGetter interface {
	Train() ITrain
}

type ITrain interface {
	Start(c context.Context, args dto.TrainStartArgs) error
	Stop(c context.Context, args dto.TrainStopArgs) error
	Log(c context.Context, args dto.TrainLogArgs) (*dto.TrainLogReply, error)
	List(c context.Context, args dto.TrainListArgs) (*dto.TrainListReply, error)
}

type train struct {
	client rest.Interface
}

func newTrain(c *APIV1Client) *train {
	return &train{
		client: c.RESTClient(),
	}
}

func (t train) Start(c context.Context, args dto.TrainStartArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("pod", "start").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t train) Stop(c context.Context, args dto.TrainStopArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("pod", "stop").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t train) Log(c context.Context, args dto.TrainLogArgs) (result *dto.TrainLogReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("pod", "log").
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

func (t train) List(c context.Context, args dto.TrainListArgs) (result *dto.TrainListReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("pod", "list").
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
