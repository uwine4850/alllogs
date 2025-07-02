package rauth

import (
	"net/http"
	"strings"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/apiform"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type MsgRegister struct {
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

func Register(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	var registerForm RegisterForm
	if err := apiform.ParseAndFill(r, &registerForm); err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}

	if strings.Trim(registerForm.Password, "") != strings.Trim(registerForm.RepeatPassword, "") {
		return api.NewClientError(http.StatusConflict, "passwords don`t match")
	}

	// Database operation.
	myauth := auth.NewAuth(w, auth.NewMysqlAuthQuery(cnf.DatabaseReader, namelib.AUTH.AUTH_TABLE), manager)
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

func createProfile(dbRead interfaces.IReadDatabase, userID int) error {
	qb := qb.NewSyncQB(dbRead.SyncQ()).Insert(cnf.DBT_PROFILE, map[string]any{"user_id": userID})
	qb.Merge()
	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
