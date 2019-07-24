package httpapi

import (
	"context"
	"net/http"
)

var _ Api = (*apiStruct)(nil)

func NewApi(b BuildRequestFunc, p ParseResponseFunc) *apiStruct {
	return &apiStruct{b, p}
}

type apiStruct struct {
	buildFn BuildRequestFunc
	parseFn ParseResponseFunc
}

func (a apiStruct) BuildRequest(ctx context.Context) (*http.Request, error) {
	return a.buildFn(ctx)
}

func (a apiStruct) ParseResponse(ctx context.Context, req *http.Request, resp *http.Response) error {
	return a.parseFn(ctx, req, resp)
}
