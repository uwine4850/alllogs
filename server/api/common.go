package api

import (
	"net/http"

	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

func SendJsonError(_error string, w http.ResponseWriter) {
	router.SendJson(map[string]string{"Error": _error}, w)
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
		SendJsonError(err.Error(), w)
	}
}
