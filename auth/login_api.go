package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/wins1908/httputil"

	"github.com/wins1908/httpapi"
)

var _ LoginApi = (*CredentialLoginApi)(nil)

type CredentialLoginApi struct {
	Endpoint   string
	Headers    map[string]string
	Credential map[string]string
	gotToken   string
}

func (a *CredentialLoginApi) BuildRequest(ctx context.Context) (*http.Request, error) {
	body, err := json.Marshal(a.Credential)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", a.Endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for header, val := range a.Headers {
		req.Header.Add(header, val)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	return req, nil
}

func (a *CredentialLoginApi) ParseResponse(ctx context.Context, req *http.Request, resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return httpapi.ErrResponseNotOk
	}
	body := new(loginApiResponseBody)
	if err := httputil.UnmarshalResponseBody(resp, body); err != nil {
		return err
	}
	if body.Token == "" {
		return httpapi.ErrNoTokenInResponseBody
	}
	a.gotToken = body.Token
	return nil
}

func (a *CredentialLoginApi) GotToken() string { return a.gotToken }

type loginApiResponseBody struct {
	Token string `json:"token"`
}
