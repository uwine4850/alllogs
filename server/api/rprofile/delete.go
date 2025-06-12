package rprofile

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Delete(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}
	profile, err := GetProfileByUID(cnf.DatabaseReader, UID.(int))
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	transaction, err := cnf.DatabaseReader.NewTransaction()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if err := transaction.BeginTransaction(); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	newQB := qb.NewSyncQB(transaction.SyncQ()).Delete(cnf.DBT_AUTH).Where(qb.Compare("id", qb.EQUAL, profile.UserId))
	newQB.Merge()
	_, err = newQB.Exec()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if err := deleteAvatar(profile); err != nil {
		if err := transaction.RollBackTransaction(); err != nil {
			return api.NewServerError(http.StatusInternalServerError, err.Error())
		}
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if err := transaction.CommitTransaction(); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
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
