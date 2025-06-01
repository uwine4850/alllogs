package mddlauth

import (
	"errors"
	"net/http"
	"slices"

	"github.com/uwine4850/alllogs/api/rauth"
	"github.com/uwine4850/foozy/pkg/builtin/builtin_mddl"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/middlewares"
)

var skipUrl = []string{"/login", "/register", "/notifications"}

func CheckJWT(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	if slices.Contains(skipUrl, r.URL.Path) {
		return nil
	}
	err := builtin_mddl.AuthJWT(
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (string, error) {
			tokenStr := r.Header.Get("Authorization")
			if tokenStr == "" {
				authJWT := r.URL.Query().Get("authJWT")
				if authJWT == "" {
					return "", errors.New("No auth JWT")
				} else {
					return authJWT, nil
				}
			}
			return tokenStr, nil
		},
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, token string, AID int) error {
			middlewares.SkipNextPage(manager.OneTimeData())
			rauth.SendLoginResponse(w, token, AID, "")
			return nil
		},
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, AID int) error {
			manager.OneTimeData().SetUserContext("AID", AID)
			return nil
		},
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
			middlewares.SkipNextPage(manager.OneTimeData())
			rauth.SendLoginResponse(w, "", 0, err.Error())
		},
	)(w, r, manager)
	return err
}
