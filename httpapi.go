package httpapi

import (
	"context"
	"errors"
	"net/http"
)

var ErrResponseNotOk = errors.New("response status is not ok")
var ErrNoDataInResponseBody = errors.New("no data in response body")
var ErrNoTokenInResponseBody = errors.New("no token in response body")

type OnFailureFunc func(ctx context.Context, req *http.Request, resp *http.Response, err error) error

type BuildRequestFunc func(ctx context.Context) (*http.Request, error)
type ParseResponseFunc func(ctx context.Context, req *http.Request, resp *http.Response) error

type Api interface {
	BuildRequest(ctx context.Context) (*http.Request, error)
	ParseResponse(ctx context.Context, req *http.Request, resp *http.Response) error
}

type Caller interface {
	Call(ctx context.Context, api Api, onFailureFn OnFailureFunc) (successful bool, err error)
}

type GetResponseBody struct {
	Data interface{} `json:"data"`
}
