package rauth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/config"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
)

func Register() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
		// Parse and validate form.
		registerForm := mydto.RegisterMessage{}
		if err := json.NewDecoder(r.Body).Decode(&registerForm); err != nil {
			return api.NewClientError(http.StatusBadRequest, err.Error())
		}
		if strings.Trim(registerForm.Password, "") != strings.Trim(registerForm.RepeatPassword, "") {
			return api.NewClientError(http.StatusConflict, "passwords don`t match")
		}

		// Database operation.
		db, err := manager.Database().ConnectionPool(config.LoadedConfig().Default.Database.MainConnectionPoolName)
		if err != nil {
			return api.NewServerError(http.StatusInternalServerError, err.Error())
		}
		myauth := auth.NewAuth(w, auth.NewMysqlAuthQuery(db, namelib.AUTH.AUTH_TABLE), manager)
		regUserId, err := myauth.RegisterUser(registerForm.Username, registerForm.Password)
		if err != nil {
			return api.NewClientError(http.StatusConflict, err.Error())
		}
		// Create profile in database.
		if err := createProfile(cnf.DatabaseReader, regUserId); err != nil {
			return api.NewServerError(http.StatusInternalServerError, err.Error())
		}
		api.SendBeseResponse(w, true, nil)
		return nil
	}
}

func createProfile(dbRead interfaces.IReadDatabase, userID int) error {
	qb := qb.NewSyncQB(dbRead.SyncQ()).Insert(cnf.DBT_PROFILE, map[string]any{"user_id": userID})
	qb.Merge()
	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
