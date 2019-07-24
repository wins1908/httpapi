package httpapi

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func IsConnectionError(err error) bool {
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return true
	}
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return true
	}
	if _, ok := err.(*net.OpError); ok {
		return true
	}
	if _, ok := err.(*url.Error); ok {
		return true
	}
	return false
}

type HttpError struct {
	Request  *http.Request
	Response *http.Response
	Err      error
}

func (e *HttpError) Error() string {
	var reqDump, respDump []byte
	if e.Request != nil {
		reqDump, _ = httputil.DumpRequest(e.Request, true)
		if e.Response != nil {
			respDump, _ = httputil.DumpResponse(e.Response, true)
		}
	}

	return fmt.Sprintf(
		"error: %s, when calling API:\n%s\nResponse:\n%s",
		e.Err.Error(),
		string(reqDump),
		string(respDump),
	)
}
