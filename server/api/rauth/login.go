package rauth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/config"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/secure"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type LoginJWTClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func Login() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
		// Parse and validate form.

		loginForm := mydto.LoginMessage{}
		if err := json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
			SendLoginResponse(w, "", 0, err.Error())
			return nil
		}

		db, err := manager.Database().ConnectionPool(config.LoadedConfig().Default.Database.MainConnectionPoolName)
		if err != nil {
			return err
		}
		myauth := auth.NewAuth(w, auth.NewMysqlAuthQuery(db, namelib.AUTH.AUTH_TABLE), manager)
		authUser, err := myauth.LoginUser(loginForm.Username, loginForm.Password)
		if err != nil {
			SendLoginResponse(w, "", 0, err.Error())
			return nil
		}

		profileId, err := getProfileIdByUserId(cnf.DatabaseReader, authUser.Id)
		if err != nil {
			SendLoginResponse(w, "", 0, err.Error())
			return nil
		}
		authClaims := auth.JWTClaims{
			Id: authUser.Id,
		}
		tokenString, err := secure.NewHmacJwtWithClaims(authClaims, manager)
		if err != nil {
			SendLoginResponse(w, "", 0, err.Error())
			return nil
		}
		SendLoginResponse(w, tokenString, profileId, "")
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
	resp := &mydto.LoginResponseMessage{
		JWT:   jwt,
		UID:   UID,
		Error: _err,
	}
	if err := mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
		api.SendJsonError(err.Error(), w)
	}
}

func getProfileIdByUserId(dbRead interfaces.IReadDatabase, id int) (int, error) {
	newQB := qb.NewSyncQB(dbRead.SyncQ()).SelectFrom("id", cnf.DBT_PROFILE).Where(qb.Compare("user_id", qb.EQUAL, id))
	newQB.Merge()
	profileId, err := newQB.Query()
	if err != nil {
		return 0, err
	}
	if len(profileId) != 1 {
		return 0, errors.New("User not found.")
	}
	intProfileId, err := dbutils.ParseInt(profileId[0]["id"])
	if err != nil {
		return 0, err
	}
	return intProfileId, nil
}
