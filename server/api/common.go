package api

import (
	"fmt"
	"net/http"

	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

func SendServerError(_error string, code int, w http.ResponseWriter) {
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
	if err := mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
		SendServerError("DTO error", http.StatusInternalServerError, w)
	}
}

func SendAnyMessage(message irest.IMessage, w http.ResponseWriter) {
	if err := mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(message)); err != nil {
		SendServerError("DTO error", http.StatusInternalServerError, w)
	}
}

func SendClientError(w http.ResponseWriter, code int, text string) {
	resp := mydto.NewClientErrorMessage(code, text)
	SendAnyMessage(resp, w)
}
