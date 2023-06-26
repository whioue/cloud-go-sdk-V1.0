package v1

import (
	"context"
	"encoding/json"
	"github.com/whioue/cloud-go-sdk-V1.0/domain/v1/task/dto"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type TaskGetter interface {
	Task() ITask
}

type ITask interface {
	Create(c context.Context, args dto.TaskCreateArgs) (*dto.TaskCreateReply, error)
	Delete(c context.Context, args dto.TaskDeleteArgs) error
	Detail(c context.Context, args dto.TaskDetailArgs) (*dto.TaskDetailReply, error)
	List(c context.Context, args dto.TaskListArgs) (*dto.TaskListReply, error)
	Update(c context.Context, args dto.TaskUpdateArgs) error
}

type task struct {
	client rest.Interface
}

func newTask(c *APIV1Client) *task {
	return &task{
		client: c.RESTClient(),
	}
}

func (t *task) Create(c context.Context, args dto.TaskCreateArgs) (result *dto.TaskCreateReply, err error) {
	data, err := t.client.
		Post().
		Resource("task").
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
func (t *task) Delete(c context.Context, args dto.TaskDeleteArgs) error {
	_, err := t.client.
		Delete().
		Resource("task").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}
	return nil
}
func (t *task) Detail(c context.Context, args dto.TaskDetailArgs) (result *dto.TaskDetailReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
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
func (t *task) List(c context.Context, args dto.TaskListArgs) (result *dto.TaskListReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("list").
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
func (t *task) Update(c context.Context, args dto.TaskUpdateArgs) error {
	_, err := t.client.
		Put().
		Resource("task").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}
