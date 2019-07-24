package httpapi

import (
	"context"
	"net/http"
)

type DoRetryFunc func(ctx context.Context) error
type RetryDecider func(ctx context.Context, req *http.Request, resp *http.Response, err error) bool

func DoRetryOnFailureFunc(fn DoRetryFunc, deciders ...RetryDecider) OnFailureFunc {
	return func(ctx context.Context, req *http.Request, resp *http.Response, err error) error {
		for _, d := range deciders {
			if !d(ctx, req, resp, err) {
				return &HttpError{Request: req, Response: resp, Err: err}
			}
		}
		return fn(ctx)
	}
}

func RetryOn5XXAndConnectionError(_ context.Context, _ *http.Request, resp *http.Response, err error) bool {
	if resp != nil && resp.StatusCode >= 500 {
		return true
	}

	if IsConnectionError(err) {
		return true
	}

	return false
}
