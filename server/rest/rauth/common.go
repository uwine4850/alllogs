package rauth

import (
	"net/http"

	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/alllogs/rest"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router/rest/restmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

func sendError(w http.ResponseWriter, _err error) func() {
	return func() {
		resp := mydto.NewBaseResponse(false, _err.Error())
		if err := restmapper.SendSafeJsonMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
			rest.SendJsonError(err.Error(), w)
		}
	}
}

func LogOut(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   namelib.AUTH.COOKIE_AUTH,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   namelib.AUTH.COOKIE_AUTH_DATE,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
