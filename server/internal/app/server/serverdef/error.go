package serverdef

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// TODO: generate all of `serverdef` package from conf / openapi declaration

// HTTPError is a custom error type that represents an error with http status code & message
type HTTPError struct {
	err  error
	code int
}

func NewHTTPError(code int, err error) *HTTPError {
	return &HTTPError{code: code, err: err}
}

func NewHTTPErrorFromString(code int, s string) *HTTPError {
	return &HTTPError{code: code, err: errors.New(s)}
}

func (h *HTTPError) Error() string {
	s := strings.Join([]string{strconv.Itoa(h.code), http.StatusText(h.code)}, " ")
	if h.err != nil {
		return fmt.Sprintf("%s:%v", s, h.err)
	}
	return s
}

func (h *HTTPError) Code() int {
	return h.code
}
