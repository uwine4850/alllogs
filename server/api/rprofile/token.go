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

func GenerateToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	genTokenRequestMessage := mydto.GenTokenMessage{}
	if err := json.NewDecoder(r.Body).Decode(&genTokenRequestMessage); err != nil {
		sendToken(w, "", err.Error())
		return nil
	}
	token, err := newToken(cnf.DatabaseReader)
	if err != nil {
		sendToken(w, "", err.Error())
		return nil
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Update(cnf.DBT_PROFILE, map[string]any{"token": token}).
		Where(qb.Compare("user_id", qb.EQUAL, genTokenRequestMessage.UserId))
	newQB.Merge()
	if _, err := newQB.Exec(); err != nil {
		sendToken(w, "", err.Error())
		return nil
	}
	sendToken(w, token, "")
	return nil
}

func DeleteToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		api.SendBeseResponse(w, false, errors.New("user ID not found"))
		return nil
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Update(cnf.DBT_PROFILE, map[string]any{"token": nil}).Where(qb.Compare("user_id", qb.EQUAL, UID))
	newQB.Merge()
	_, err := newQB.Exec()
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
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
	if err := mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(resp)); err != nil {
		api.SendServerError("DTO error", http.StatusInternalServerError, w)
	}
}
