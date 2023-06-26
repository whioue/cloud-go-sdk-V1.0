package v1

import (
	"context"
	"github.com/whioue/cloud-go-sdk-V1.0/domain/v1/log/dto"
	"github.com/whioue/cloud-go-sdk-V1.0/pkg/json"
	"github.com/whioue/cloud-go-sdk-V1.0/rest"
)

type LogGetter interface {
	Log() ILog
}

type ILog interface {
	Delete(c context.Context, args dto.LogDeleteArgs) error
	Download(c context.Context, args dto.LogDownloadArgs) (*dto.LogDownloadReply, error)
	List(c context.Context, args dto.LogListArgs) (*dto.LogListReply, error)
	Detail(c context.Context, args dto.LogDetailArgs) (*dto.LogDetailReply, error)
}

type log struct {
	client rest.Interface
}

func (l *log) Delete(c context.Context, args dto.LogDeleteArgs) error {
	return nil
}

func (l *log) Download(c context.Context, args dto.LogDownloadArgs) (*dto.LogDownloadReply, error) {
	return nil, nil
}

func (l *log) List(c context.Context, args dto.LogListArgs) (result *dto.LogListReply, err error) {

	data, err := l.client.
		Get().
		Resource("log").
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

func (l *log) Detail(c context.Context, args dto.LogDetailArgs) (*dto.LogDetailReply, error) {
	return nil, nil
}

func newLog(c *APIV1Client) *log {
	return &log{
		client: c.RESTClient(),
	}
}
