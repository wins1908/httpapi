package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wins1908/httpapi"
)

var _ httpapi.Api = &BearerHeaderApi{}

type BearerHeaderApi struct {
	Api   httpapi.Api
	Token string
}

func (a BearerHeaderApi) BuildRequest(ctx context.Context) (*http.Request, error) {
	req, err := a.Api.BuildRequest(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Token))
	return req, nil
}

func (a BearerHeaderApi) ParseResponse(ctx context.Context, req *http.Request, resp *http.Response) error {
	if resp.StatusCode == http.StatusUnauthorized {
		return ErrResponseUnauthorized
	}

	return a.Api.ParseResponse(ctx, req, resp)
}
