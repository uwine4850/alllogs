package rprofile

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

func GenerateToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	genTokenRequestMessage := mydto.GenTokenMessage{}
	if err := json.NewDecoder(r.Body).Decode(&genTokenRequestMessage); err != nil {
		return sendToken(w, "", err.Error())
	}
	token, err := newToken(cnf.DatabaseReader)
	if err != nil {
		return sendToken(w, "", err.Error())
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Update(cnf.DBT_PROFILE, map[string]any{"token": token}).Where(qb.Compare("id", qb.EQUAL, genTokenRequestMessage.UserId))
	newQB.Merge()
	if _, err := newQB.Exec(); err != nil {
		return sendToken(w, "", err.Error())
	}
	return func() {
		sendToken(w, token, "")()
	}
}

func DeleteToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		return api.SendBeseResponse(w, false, errors.New("user ID not found"))
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Update(cnf.DBT_PROFILE, map[string]any{"token": nil}).Where(qb.Compare("user_id", qb.EQUAL, AID))
	newQB.Merge()
	_, err := newQB.Exec()
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	return api.SendBeseResponse(w, true, nil)
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

func sendToken(w http.ResponseWriter, token string, _err string) func() {
	resp := &mydto.TokenResponse{
		Token: token,
		Error: _err,
	}
	return func() {
		if err := mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
			api.SendJsonError(err.Error(), w)
		}
	}
}
