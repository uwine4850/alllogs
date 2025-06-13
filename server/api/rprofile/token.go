package rprofile

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type TokenForm struct {
	UserId int `form:"UserId" empty:"-err"`
}

func GenerateToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	var tokenForm TokenForm
	if err := mapper.FillStructFromForm(frm, &tokenForm); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	token, err := newToken(cnf.DatabaseReader)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Update(cnf.DBT_PROFILE, map[string]any{"token": token}).
		Where(qb.Compare("user_id", qb.EQUAL, tokenForm.UserId))
	newQB.Merge()
	if _, err := newQB.Exec(); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	sendToken(w, token, "")
	return nil
}

func DeleteToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Update(cnf.DBT_PROFILE, map[string]any{"token": nil}).Where(qb.Compare("user_id", qb.EQUAL, UID))
	newQB.Merge()
	_, err := newQB.Exec()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}

func newToken(dbRead interfaces.IReadDatabase) (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	token := hex.EncodeToString(bytes)
	exist, err := tokenExists(dbRead, token)
	if err != nil {
		return "", err
	}
	if exist {
		nToken, err := newToken(dbRead)
		if err != nil {
			return "", err
		}
		return nToken, nil
	}
	return token, nil
}

func tokenExists(dbRead interfaces.IReadDatabase, token string) (bool, error) {
	newQB := qb.NewSyncQB(dbRead.SyncQ()).
		Select(qb.Exists(qb.SQ(
			false,
			qb.NewNoDbQB().SelectFrom(1, cnf.DBT_PROFILE).Where(qb.Compare("token", qb.EQUAL, token)),
		))).As("exist")
	newQB.Merge()
	exists, err := newQB.Query()
	if err != nil {
		return false, err
	}
	existInt, err := dbutils.ParseInt(exists[0]["exist"])
	if err != nil {
		return false, err
	}
	return existInt != 0, nil
}

func sendToken(w http.ResponseWriter, token string, _err string) {
	resp := &mydto.TokenResponse{
		Token: token,
		Error: _err,
	}
	if err := mapper.SendSafeJsonDTOMessage(w, http.StatusOK, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
		api.SendServerError(w, http.StatusInternalServerError, "DTO error")
	}
}
