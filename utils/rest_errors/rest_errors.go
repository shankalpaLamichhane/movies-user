package rest_errors

import "net/http"

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Message() string{
	return e.ErrMessage
}

func (e restErr) Status() int{
	return e.ErrStatus
}

func (e restErr) Error() string{
	return e.ErrError
}

func (e restErr) Causes() []interface{}{
	return e.ErrCauses
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
		ErrCauses:  nil,
	}
}
