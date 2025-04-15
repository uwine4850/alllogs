package mddlauth

import (
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/rest/rauth"
	"github.com/uwine4850/foozy/pkg/builtin/builtin_mddl"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/middlewares"
)

func CheckJWT(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) {
	if r.URL.Path != "/isauth" {
		return
	}
	builtin_mddl.AuthJWT(
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (string, error) {
			tokenStr := r.Header.Get("Authorization")
			if tokenStr == "" {
				return "", errors.New("No auth JWT")
			}
			return tokenStr, nil
		},
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, token string, UID string) error {
			middlewares.SkipNextPage(manager.OneTimeData())
			rauth.SendLoginResponse(w, token, UID, "")()
			return nil
		},
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
			middlewares.SkipNextPage(manager.OneTimeData())
			rauth.SendLoginResponse(w, "", "", err.Error())()
		},
	)(w, r, manager)
}
