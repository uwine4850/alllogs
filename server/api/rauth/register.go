package rauth

import (
	"net/http"
	"strings"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/config"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type RegisterMessage struct {
	rest.ImplementDTOMessage
	TypRegisterMessage rest.TypeId `dto:"-typeid"`
	Username           string      `dto:"Username"`
	Password           string      `dto:"Password"`
	RepeatPassword     string      `dto:"RepeatPassword"`
}

type RegisterForm struct {
	Username       string `form:"Username" emty:"-err"`
	Password       string `form:"Password" emty:"-err"`
	RepeatPassword string `form:"RepeatPassword" emty:"-err"`
}

func Register() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
		frm := form.NewForm(r)
		if err := frm.Parse(); err != nil {
			return api.NewClientError(http.StatusBadRequest, err.Error())
		}
		var registerForm RegisterForm
		if err := mapper.FillStructFromForm(frm, &registerForm); err != nil {
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
