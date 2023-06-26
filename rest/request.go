package rest

import (
	"context"
	"errors"
	"fmt"
	http2 "github.com/whioue/cloud-go-sdk-V1.0/pkg/http"
	"net/http"

	"net/url"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/whioue/cloud-go-sdk-V1.0/pkg/runtime"
	"github.com/whioue/cloud-go-sdk-V1.0/third_party/forked/gorequest"
)

type Request struct {
	c *RESTClient

	timeout    time.Duration
	verb       string
	pathPrefix string
	subpath    string
	params     url.Values
	headers    http.Header

	resource     string
	resourceName string
	subresource  string

	// output
	err  error
	body interface{}
}

func NewRequest(c *RESTClient) *Request {
	var pathPrefix string
	if c.base != nil {
		pathPrefix = path.Join("/", c.base.Path, c.versionedAPIPath)
	} else {
		pathPrefix = path.Join("/", c.versionedAPIPath)
	}

	r := &Request{
		c:          c,
		pathPrefix: pathPrefix,
	}

	switch {
	case len(c.content.AcceptContentTypes) > 0:
		r.SetHeader("Accept", c.content.AcceptContentTypes)
	case len(c.content.ContentType) > 0:
		r.SetHeader("Accept", c.content.ContentType+", */*")
	}

	return r
}

func NewRequestWithClient(base *url.URL, versionedAPIPath string,
	content ClientContentConfig, client *gorequest.SuperAgent) *Request {
	return NewRequest(&RESTClient{
		base:             base,
		versionedAPIPath: versionedAPIPath,
		content:          content,
		Client:           client,
	})
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

func (r *Request) Prefix(segments ...string) *Request {
	if r.err != nil {
		return r
	}

	r.pathPrefix = path.Join(r.pathPrefix, path.Join(segments...))

	return r
}

func (r *Request) Suffix(segments ...string) *Request {
	if r.err != nil {
		return r
	}

	r.subpath = path.Join(r.subpath, path.Join(segments...))

	return r
}

func (r *Request) Resource(resource string) *Request {
	if r.err != nil {
		return r
	}

	if len(r.resource) != 0 {
		r.err = fmt.Errorf("resource already set to %q, cannot change to %q", r.resource, resource)
		return r
	}

	if msgs := IsValidPathSegmentName(resource); len(msgs) != 0 {
		r.err = fmt.Errorf("invalid resource %q: %v", resource, msgs)
		return r
	}

	r.resource = resource

	return r
}

func (r *Request) SubResource(subresources ...string) *Request {
	if r.err != nil {
		return r
	}

	subresource := path.Join(subresources...)

	if len(r.subresource) != 0 {
		r.err = fmt.Errorf("subresource already set to %q, cannot change to %q", r.resource, subresource)
		return r
	}

	for _, s := range subresources {
		if msgs := IsValidPathSegmentName(s); len(msgs) != 0 {
			r.err = fmt.Errorf("invalid subresource %q: %v", s, msgs)
			return r
		}
	}

	r.subresource = subresource

	return r
}

func (r *Request) Name(resourceName string) *Request {
	if r.err != nil {
		return r
	}

	if len(resourceName) == 0 {
		r.err = fmt.Errorf("resource name may not be empty")
		return r
	}

	if len(r.resourceName) != 0 {
		r.err = fmt.Errorf("resource name already set to %q, cannot change to %q", r.resourceName, resourceName)
		return r
	}

	if msgs := IsValidPathSegmentName(resourceName); len(msgs) != 0 {
		r.err = fmt.Errorf("invalid resource name %q: %v", resourceName, msgs)
		return r
	}

	r.resourceName = resourceName

	return r
}

func (r *Request) AbsPath(segments ...string) *Request {
	if r.err != nil {
		return r
	}

	r.pathPrefix = path.Join(r.c.base.Path, path.Join(segments...))

	if len(segments) == 1 && (len(r.c.base.Path) > 1 || len(segments[0]) > 1) && strings.HasSuffix(segments[0], "/") {
		// preserve any trailing slashes for legacy behavior
		r.pathPrefix += "/"
	}

	return r
}

func (r *Request) RequestURI(uri string) *Request {
	if r.err != nil {
		return r
	}

	locator, err := url.Parse(uri)
	if err != nil {
		r.err = err
		return r
	}

	r.pathPrefix = locator.Path

	if len(locator.Query()) > 0 {
		if r.params == nil {
			r.params = make(url.Values)
		}

		for k, v := range locator.Query() {
			r.params[k] = v
		}
	}

	return r
}

func (r *Request) Param(paramName, s string) *Request {
	if r.err != nil {
		return r
	}

	return r.setParam(paramName, s)
}

func (r *Request) VersionedParams(v interface{}) *Request {
	if r.err != nil {
		return r
	}

	r.c.Client.Query(v)

	return r
}

func (r *Request) setParam(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}

	r.params[paramName] = append(r.params[paramName], value)

	return r
}

func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}

	r.headers.Del(key)

	for _, value := range values {
		r.headers.Add(key, value)
	}

	return r
}

func (r *Request) Timeout(d time.Duration) *Request {
	if r.err != nil {
		return r
	}

	r.timeout = d

	return r
}

func (r *Request) URL() *url.URL {
	p := r.pathPrefix
	if len(r.resource) != 0 {
		p = path.Join(p, strings.ToLower(r.resource))
	}
	if len(r.resourceName) != 0 || len(r.subpath) != 0 || len(r.subresource) != 0 {
		p = path.Join(p, r.resourceName, r.subresource, r.subpath)
	}

	finalURL := &url.URL{}
	if r.c.base != nil {
		*finalURL = *r.c.base
	}

	finalURL.Path = p

	query := url.Values{}

	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	if r.timeout != 0 {
		query.Set("timeout", r.timeout.String())
	}

	finalURL.RawQuery = query.Encode()

	return finalURL
}

func (r *Request) Body(obj interface{}) *Request {
	if v := reflect.ValueOf(obj); v.Kind() == reflect.Struct {
		r.SetHeader("Content-Type", r.c.content.ContentType)
	}

	r.body = obj

	return r
}

func (r *Request) Do(ctx context.Context) Result {
	client := r.c.Client
	client.Header = r.headers

	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.timeout)

		defer cancel()
	}

	client.WithContext(ctx)

	resp, body, errs := client.CustomMethod(r.verb, r.URL().String()).Send(r.body).EndBytes()
	if err := combineErr(resp, body, errs); err != nil {
		return Result{
			response: &resp,
			err:      err,
			body:     body,
		}
	}

	decoder, err := r.c.content.Negotiator.Decoder()
	encoder, err := r.c.content.Negotiator.Encoder()
	if err != nil {
		return Result{
			response: &resp,
			err:      err,
			body:     body,
			decoder:  decoder,
			encoder:  encoder,
		}
	}

	return Result{
		response: &resp,
		body:     body,
		decoder:  decoder,
		encoder:  encoder,
	}
}

type Result struct {
	response *gorequest.Response
	err      error
	body     []byte
	decoder  runtime.Decoder
	encoder  runtime.Encoder
}

func (r Result) Raw() ([]byte, error) {
	return r.body, r.err
}

func (r Result) Data() ([]byte, error) {
	if r.err != nil {
		return []byte{}, r.Error()
	}

	if r.decoder == nil {
		return []byte{}, fmt.Errorf("serializer doesn't exist")
	}

	httpData, err := r.parseHttpResp()
	if err != nil {
		return []byte{}, err
	}
	return r.encoder.Encode(httpData)
}

func (r Result) Error() error {
	return r.err
}

func (r Result) parseHttpResp() (interface{}, error) {
	if r.err != nil {
		return nil, r.Error()
	}

	if r.decoder == nil {
		return nil, fmt.Errorf("serializer doesn't exist")
	}
	httpResp := &http2.Resp{}
	err := r.decoder.Decode(r.body, httpResp)
	if err != nil {
		return nil, err
	}
	return httpResp.Data, nil
}

func combineErr(resp gorequest.Response, body []byte, errs []error) error {
	var e, sep string

	if len(errs) > 0 {
		for _, err := range errs {
			e = sep + err.Error()
			sep = "\n"
		}

		return errors.New(e)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	return nil
}

var NameMayNotBe = []string{".", ".."}

var NameMayNotContain = []string{"/", "%"}

func IsValidPathSegmentName(name string) []string {
	for _, illegalName := range NameMayNotBe {
		if name == illegalName {
			return []string{fmt.Sprintf(`may not be '%s'`, illegalName)}
		}
	}

	var errors []string

	for _, illegalContent := range NameMayNotContain {
		if strings.Contains(name, illegalContent) {
			errors = append(errors, fmt.Sprintf(`may not contain '%s'`, illegalContent))
		}
	}

	return errors
}

func IsValidPathSegmentPrefix(name string) []string {
	var errors []string

	for _, illegalContent := range NameMayNotContain {
		if strings.Contains(name, illegalContent) {
			errors = append(errors, fmt.Sprintf(`may not contain '%s'`, illegalContent))
		}
	}

	return errors
}

// ValidatePathSegmentName validates the name can be safely encoded as a path segment.
func ValidatePathSegmentName(name string, prefix bool) []string {
	if prefix {
		return IsValidPathSegmentPrefix(name)
	}

	return IsValidPathSegmentName(name)
}
