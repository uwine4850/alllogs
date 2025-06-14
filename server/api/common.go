package api

import (
	"fmt"
	"net/http"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/rest"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type BaseResponseMessage struct {
	rest.ImplementDTOMessage
	TypBaseResponseMessage rest.TypeId `dto:"-typeid"`
	Ok                     bool        `dto:"Ok"`
	Error                  string      `dto:"Error"`
}

func NewBaseResponse(ok bool, error string) *BaseResponseMessage {
	return &BaseResponseMessage{
		Ok:    ok,
		Error: error,
	}
}

// 400 - 499
type ClientErrorMessage struct {
	rest.ImplementDTOMessage
	TypClientErrorMessage rest.TypeId `dto:"-typeid"`
	Code                  int         `dto:"Code"`
	Text                  string      `dto:"Text"`
}

func NewClientErrorMessage(code int, text string) *ClientErrorMessage {
	return &ClientErrorMessage{
		Code: code,
		Text: text,
	}
}

// 500 - 511
type ServerErrorMessage struct {
	rest.ImplementDTOMessage
	TypServerErrorMessage rest.TypeId `dto:"-typeid"`
	Code                  int         `dto:"Code"`
	Text                  string      `dto:"Text"`
}

func NewServerErrorMessage(code int, text string) *ServerErrorMessage {
	return &ServerErrorMessage{
		Code: code,
		Text: text,
	}
}

func sendUntypeServerError(_error string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	fmt.Fprint(w, _error)
}

func SendBeseResponse(w http.ResponseWriter, ok bool, _err error) {
	var errValue string
	if _err != nil {
		errValue = _err.Error()
	} else {
		errValue = ""
	}
	resp := NewBaseResponse(ok, errValue)
	if err := mapper.SendSafeJsonDTOMessage(w, http.StatusOK, cnf.DTO, typeopr.Ptr{}.New(resp)); err != nil {
		sendUntypeServerError("DTO error", http.StatusInternalServerError, w)
	}
}

func SendAnyMessage(message irest.IMessage, w http.ResponseWriter, code int) {
	if err := mapper.SendSafeJsonDTOMessage(w, code, cnf.DTO, typeopr.Ptr{}.New(message)); err != nil {
		sendUntypeServerError("DTO error", http.StatusInternalServerError, w)
	}
}

func SendClientError(w http.ResponseWriter, code int, text string) {
	resp := NewClientErrorMessage(code, text)
	SendAnyMessage(resp, w, code)
}

func SendServerError(w http.ResponseWriter, code int, text string) {
	resp := NewServerErrorMessage(code, text)
	SendAnyMessage(resp, w, code)
}

type ClientError struct {
	Code int
	Text string
}

func NewClientError(code int, text string) *ClientError {
	return &ClientError{
		Code: code,
		Text: text,
	}
}

func (ce *ClientError) Error() string {
	return ce.Text
}

type ServerError struct {
	Code int
	Text string
}

func NewServerError(code int, text string) *ServerError {
	return &ServerError{
		Code: code,
		Text: text,
	}
}

func (ce *ServerError) Error() string {
	return ce.Text
}
