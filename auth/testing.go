package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestdataLoginResponseOk(token string) *http.Response {
	body := &loginApiResponseBody{Token: token}
	b, _ := json.Marshal(body)
	reader := bytes.NewReader(b)

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(reader),
	}
}

func AssertLoginRequest(
	t assert.TestingT,
	expectUrl string,
	expectHeaders map[string]string,
	expectCredentials map[string]string,
	actual *http.Request,
) {
	assert.Equal(t, http.MethodPost, actual.Method)
	assert.Equal(t, "application/json", actual.Header.Get("Accept"))
	for header, val := range expectHeaders {
		assert.Equal(t, val, actual.Header.Get(header))
	}

	assert.Equal(t, expectUrl, actual.URL.String())

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(actual.Body); err != nil {
		assert.Fail(t, "cannot read from actual request body")
	}
	expectBody, _ := json.Marshal(expectCredentials)
	assert.JSONEq(t, string(expectBody), buf.String())
}
