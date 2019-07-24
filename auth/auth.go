package auth

import (
	"context"
	"errors"

	"github.com/wins1908/httpapi"
)

var ErrResponseUnauthorized = errors.New("unauthorized")

type LoginApi interface {
	httpapi.Api
	GotToken() string
}

type LoginAuthorizer interface {
	AcquireToken(ctx context.Context, gotTokenFn GotTokenFunc, failure httpapi.OnFailureFunc) error
	ClearToken(ctx context.Context) error
}

type GotTokenFunc func(ctx context.Context, token string, loginApiCalled bool) error
