package rauth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/rest/restmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

func Register() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
		// Parse and validate form.
		registerForm := mydto.RegisterMessage{}
		if err := json.NewDecoder(r.Body).Decode(&registerForm); err != nil {
			return api.SendBeseResponse(w, false, err)
		}
		if strings.Trim(registerForm.Password, "") != strings.Trim(registerForm.RepeatPassword, "") {
			return api.SendBeseResponse(w, false, errors.New("passwords don`t match"))
		}

		// Database operation.
		db := database.NewDatabase(cnf.DATABASE_ARGS)
		if err := db.Connect(); err != nil {
			return api.SendBeseResponse(w, false, err)
		}
		defer func() {
			if err := db.Close(); err != nil {
				api.SendBeseResponse(w, false, err)()
			}
		}()
		myauth := auth.NewAuth(db, w, manager)
		regUserId, err := myauth.RegisterUser(registerForm.Username, registerForm.Password)
		if err != nil {
			return api.SendBeseResponse(w, false, err)
		}
		// Create profile in database.
		if err := createProfile(db, regUserId); err != nil {
			return api.SendBeseResponse(w, false, err)
		}
		return func() {
			resp := mydto.NewBaseResponse(true, "")
			if err := restmapper.SendSafeJsonMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
				api.SendJsonError(err.Error(), w)
			}
		}
	}
}

func createProfile(db *database.Database, userID int) error {
	qb := qb.NewSyncQB(db.SyncQ()).Insert(cnf.DBT_PROFILE, map[string]any{"user_id": userID})
	qb.Merge()
	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
