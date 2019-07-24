package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/mock"
)

// MockApi is an autogenerated mock type for the Api type
type MockApi struct {
	mock.Mock
}

// BuildRequest provides a mock function with given fields: ctx
func (_m *MockApi) BuildRequest(ctx context.Context) (*http.Request, error) {
	ret := _m.Called(ctx)

	var r0 *http.Request
	if rf, ok := ret.Get(0).(func(context.Context) *http.Request); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseResponse provides a mock function with given fields: ctx, req, resp
func (_m *MockApi) ParseResponse(ctx context.Context, req *http.Request, resp *http.Response) error {
	ret := _m.Called(ctx, req, resp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request, *http.Response) error); ok {
		r0 = rf(ctx, req, resp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCaller is an autogenerated mock type for the Caller type
type MockCaller struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, api, onFailureFn
func (_m *MockCaller) Call(ctx context.Context, api Api, onFailureFn OnFailureFunc) (bool, error) {
	ret := _m.Called(ctx, api, onFailureFn)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, Api, OnFailureFunc) bool); ok {
		r0 = rf(ctx, api, onFailureFn)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, Api, OnFailureFunc) error); ok {
		r1 = rf(ctx, api, onFailureFn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRetryHandler is an autogenerated mock type for the RetryHandler type
type MockRetryHandler struct {
	mock.Mock
}

// BuildOnFailureFunc provides a mock function with given fields: fn, deciders
func (_m *MockRetryHandler) BuildOnFailureFunc(fn DoRetryFunc, deciders ...RetryDecider) OnFailureFunc {
	_va := make([]interface{}, len(deciders))
	for _i := range deciders {
		_va[_i] = deciders[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fn)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 OnFailureFunc
	if rf, ok := ret.Get(0).(func(DoRetryFunc, ...RetryDecider) OnFailureFunc); ok {
		r0 = rf(fn, deciders...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(OnFailureFunc)
		}
	}

	return r0
}

func TestdataGetResponseOk(data interface{}) *http.Response {
	body := &GetResponseBody{Data: data}
	b, _ := json.Marshal(body)
	reader := bytes.NewReader(b)

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(reader),
	}
}