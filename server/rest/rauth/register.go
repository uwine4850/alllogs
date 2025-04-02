package rauth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/alllogs/rest"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest/restmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

func Register() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
		// Parse and validate form.
		frm := form.NewForm(r)
		if err := frm.Parse(); err != nil {
			return sendError(w, err)
		}
		registerForm := mydto.Register{}
		if err := json.NewDecoder(r.Body).Decode(&registerForm); err != nil {
			return sendError(w, err)
		}
		if strings.Trim(registerForm.Password, "") != strings.Trim(registerForm.RepeatPassword, "") {
			return sendError(w, errors.New("passwords don`t match"))
		}

		// Database operation.
		db := database.NewDatabase(cnf.DATABASE_ARGS)
		if err := db.Connect(); err != nil {
			return sendError(w, err)
		}
		defer func() {
			if err := db.Close(); err != nil {
				sendError(w, err)
			}
		}()
		myauth := auth.NewAuth(db, w, manager)
		if err := myauth.RegisterUser(registerForm.Username, registerForm.Password); err != nil {
			return sendError(w, err)
		}

		return func() {
			resp := mydto.NewBaseResponse(true, "")
			if err := restmapper.SendSafeJsonMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
				rest.SendJsonError(err.Error(), w)
			}
		}
	}
}

func sendError(w http.ResponseWriter, _err error) func() {
	return func() {
		resp := mydto.NewBaseResponse(false, _err.Error())
		if err := restmapper.SendSafeJsonMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
			rest.SendJsonError(err.Error(), w)
		}
	}
}
