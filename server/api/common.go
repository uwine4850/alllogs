package api

import (
	"fmt"
	"net/http"

	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

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
	resp := mydto.NewBaseResponse(ok, errValue)
	if err := mapper.SendSafeJsonDTOMessage(w, http.StatusOK, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
		sendUntypeServerError("DTO error", http.StatusInternalServerError, w)
	}
}

func SendAnyMessage(message irest.IMessage, w http.ResponseWriter, code int) {
	if err := mapper.SendSafeJsonDTOMessage(w, code, mydto.DTO, typeopr.Ptr{}.New(message)); err != nil {
		sendUntypeServerError("DTO error", http.StatusInternalServerError, w)
	}
}

func SendClientError(w http.ResponseWriter, code int, text string) {
	resp := mydto.NewClientErrorMessage(code, text)
	SendAnyMessage(resp, w, code)
}

func SendServerError(w http.ResponseWriter, code int, text string) {
	resp := mydto.NewServerErrorMessage(code, text)
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
