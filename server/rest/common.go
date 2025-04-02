package rest

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/router"
)

func SendJsonError(_error string, w http.ResponseWriter) {
	router.SendJson(map[string]string{"Error": _error}, w)
}
