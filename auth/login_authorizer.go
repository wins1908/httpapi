package auth

import (
	"context"
	"fmt"

	"github.com/wins1908/strcache"

	"github.com/wins1908/httpapi"
)

type BuildLoginApiFunc func(ctx context.Context) LoginApi

func NewLoginAuthorizer(
	c httpapi.Caller,
	buildLoginApiFn BuildLoginApiFunc,
	tokenFetcher strcache.Fetcher,
	tokenKey string,
) *loginAuthorizer {
	return &loginAuthorizer{c, buildLoginApiFn, tokenFetcher, tokenKey}
}

type loginAuthorizer struct {
	caller          httpapi.Caller
	buildLoginApiFn BuildLoginApiFunc
	tokenFetcher    strcache.Fetcher
	tokenKey        string
}

func (a *loginAuthorizer) AcquireToken(ctx context.Context, gotTokenFn GotTokenFunc, failure httpapi.OnFailureFunc) error {
	return a.tokenFetcher.Fetch(
		ctx,
		a.tokenKey,
		func(ctx context.Context, newFn strcache.NewValueFunc) error {
			loginApi := a.buildLoginApiFn(ctx)
			loginSuccess, err := a.caller.Call(ctx, loginApi, failure)
			if loginSuccess {
				return newFn(ctx, loginApi.GotToken())
			}
			return err
		}, func(ctx context.Context, token string, isNew bool) error {
			return gotTokenFn(ctx, token, isNew)
		},
	)
}

func (a *loginAuthorizer) ClearToken(ctx context.Context) error {
	err := a.tokenFetcher.Clear(ctx, a.tokenKey)
	return err
}

type LoginAuthorizerMap map[string]LoginAuthorizer

func (m LoginAuthorizerMap) AcquireToken(
	ctx context.Context,
	mapKey string,
	gotTokenFn GotTokenFunc,
	failureFn httpapi.OnFailureFunc,
) error {
	if a, exists := m[mapKey]; exists {
		return a.AcquireToken(ctx, gotTokenFn, failureFn)
	}

	return fmt.Errorf("no authorizer on map key `%s`", mapKey)
}

func (m LoginAuthorizerMap) ClearToken(ctx context.Context, mapKey string) error {
	if a, exists := m[mapKey]; exists {
		return a.ClearToken(ctx)
	}
	return fmt.Errorf("no authorizer for map key `%s`", mapKey)
}
