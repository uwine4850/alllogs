package rprofile

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Delete(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		api.SendBeseResponse(w, false, errors.New("user ID not found"))
		return nil
	}
	profile, err := GetProfileByAID(cnf.DatabaseReader, AID.(int))
	if err != nil {
		panic(err)
	}
	transaction := cnf.DatabaseReader.NewTransaction()
	newQB := qb.NewSyncQB(transaction.SyncQ()).Delete(cnf.DBT_AUTH).Where(qb.Compare("id", qb.EQUAL, profile.UserId))
	newQB.Merge()
	_, err = newQB.Exec()
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	if err := deleteAvatar(profile); err != nil {
		if err := transaction.RollBackTransaction(); err != nil {
			api.SendBeseResponse(w, false, err)
			return nil
		}
		api.SendBeseResponse(w, false, err)
		return nil
	}
	if err := transaction.CommitTransaction(); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}

func deleteAvatar(profile *ProfileDBView) error {
	if profile.Avatar != cnf.DEFAULT_AVATAR_PATH {
		relPath := filepath.Join("../client/public/", profile.Avatar)
		err := os.Remove(relPath)
		if err != nil {
			return err
		}
	}
	return nil
}
