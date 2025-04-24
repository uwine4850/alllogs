package rauth

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/namelib"
)

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
