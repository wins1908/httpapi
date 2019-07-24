package httpapi

import (
	"context"
	"net/http"
)

var _ Caller = &caller{}

func NewCaller(c *http.Client) *caller {
	return &caller{c}
}

type caller struct {
	httpClient *http.Client
}

func (c caller) Call(ctx context.Context, api Api, onFailureFn OnFailureFunc) (successful bool, err error) {
	req, err := api.BuildRequest(ctx)
	if err != nil {
		return false, onFailureFn(ctx, req, nil, err)
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, onFailureFn(ctx, req, resp, err)
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	if err := api.ParseResponse(ctx, req, resp); err != nil {
		return false, onFailureFn(ctx, req, resp, err)
	}

	return true, nil
}
