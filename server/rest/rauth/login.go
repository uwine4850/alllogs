package rauth

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/alllogs/rest"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest/restmapper"
	"github.com/uwine4850/foozy/pkg/secure"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type LoginJWTClaims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func Login() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
		// Parse and validate form.
		frm := form.NewForm(r)
		if err := frm.Parse(); err != nil {
			return SendLoginResponse(w, "", err.Error())
		}

		loginForm := mydto.Login{}
		if err := json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
			return SendLoginResponse(w, "", err.Error())
		}

		// Database operation.
		db := database.NewDatabase(cnf.DATABASE_ARGS)
		if err := db.Connect(); err != nil {
			return SendLoginResponse(w, "", err.Error())
		}
		defer func() {
			if err := db.Close(); err != nil {
				SendLoginResponse(w, "", err.Error())
			}
		}()
		// userDB, err := auth.UserByUsername(db, loginForm.Username)
		// if err != nil {
		// 	return SendLoginResponse(w, "", err.Error())
		// }
		// if userDB == nil {
		// 	return sendError(w, &auth.ErrUserNotExist{Username: loginForm.Username})
		// }
		// err = auth.ComparePassword(dbutils.ParseString(userDB["password"]), loginForm.Password)
		// if err != nil {
		// 	return SendLoginResponse(w, "", err.Error())
		// }
		// var authUser auth.AuthUser
		// if err := dbmapper.FillStructFromDb(userDB, typeopr.Ptr{}.New(&authUser)); err != nil {
		// 	return SendLoginResponse(w, "", err.Error())
		// }
		// // JWT
		// tokenString, err := NewLoginJWT(authUser.Id, manager)
		// if err != nil {
		// 	return SendLoginResponse(w, "", err.Error())
		// }
		myauth := auth.NewAuth(db, w, manager)
		authUser, err := myauth.LoginUser(loginForm.Username, loginForm.Password)
		if err != nil {
			return SendLoginResponse(w, "", err.Error())
		}
		authClaims := auth.JWTClaims{
			Id: authUser.Id,
		}
		tokenString, err := secure.NewHmacJwtWithClaims(authClaims, manager)
		if err != nil {
			return SendLoginResponse(w, "", err.Error())
		}
		return SendLoginResponse(w, tokenString, "")
	}
}

func NewLoginJWT(uid string, manager interfaces.IManager) (string, error) {
	claims := &LoginJWTClaims{
		Id: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(manager.Key().Get32BytesKey().HashKey()))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func SendLoginResponse(w http.ResponseWriter, jwt string, _err string) func() {
	return func() {
		resp := &mydto.LoginResponse{
			JWT:   jwt,
			Error: _err,
		}
		if err := restmapper.SendSafeJsonMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
			rest.SendJsonError(err.Error(), w)
		}
	}
}
