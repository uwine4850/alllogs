package rauth

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/config"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest"
	"github.com/uwine4850/foozy/pkg/secure"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type MsgLogin struct {
	rest.ImplementDTOMessage
	TypLoginMessage rest.TypeId `dto:"-typeid"`
	Username        string      `dto:"Username"`
	Password        string      `dto:"Password"`
}

type MsgLoginResponse struct {
	rest.ImplementDTOMessage
	TypLoginResponseMessage rest.TypeId `dto:"-typeid"`
	JWT                     string      `dto:"JWT"`
	UID                     int         `dto:"UID"`
	Error                   string      `dto:"Error"`
}

type LoginJWTClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

type LoginForm struct {
	Username string `form:"Username" empty:"-err"`
	Password string `form:"Password" empty:"-err"`
}

func Login() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
		frm := form.NewForm(r)
		if err := frm.Parse(); err != nil {
			return api.NewServerError(http.StatusBadRequest, err.Error())
		}
		var loginForm LoginForm
		if err := mapper.FillStructFromForm(frm, &loginForm); err != nil {
			return api.NewServerError(http.StatusBadRequest, err.Error())
		}

		db, err := manager.Database().ConnectionPool(config.LoadedConfig().Default.Database.MainConnectionPoolName)
		if err != nil {
			return api.NewServerError(http.StatusInternalServerError, err.Error())
		}
		myauth := auth.NewAuth(w, auth.NewMysqlAuthQuery(db, namelib.AUTH.AUTH_TABLE), manager)
		authUser, err := myauth.LoginUser(loginForm.Username, loginForm.Password)
		if err != nil {
			if errors.As(err, &auth.ErrUserNotExist{}) || errors.As(err, &auth.ErrPasswordsDontMatch{}) {
				return api.NewClientError(http.StatusConflict, err.Error())
			} else {
				return api.NewServerError(http.StatusInternalServerError, err.Error())
			}
		}

		authClaims := auth.JWTClaims{
			Id: authUser.Id,
		}
		tokenString, err := secure.NewHmacJwtWithClaims(authClaims, manager)
		if err != nil {
			return api.NewServerError(http.StatusInternalServerError, err.Error())
		}
		SendLoginResponse(w, tokenString, authUser.Id, "")
		return nil
	}
}

func NewLoginJWT(uid int, manager interfaces.IManager) (string, error) {
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

func SendLoginResponse(w http.ResponseWriter, jwt string, UID int, _err string) {
	resp := &MsgLoginResponse{
		JWT:   jwt,
		UID:   UID,
		Error: _err,
	}
	if err := mapper.SendSafeJsonDTOMessage(w, http.StatusOK, cnf.DTO, typeopr.Ptr{}.New(resp)); err != nil {
		api.SendServerError(w, http.StatusInternalServerError, "DTO error")
	}
}
