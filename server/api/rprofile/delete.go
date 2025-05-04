package rprofile

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Delete(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		return api.SendBeseResponse(w, false, errors.New("user ID not found"))
	}
	db := database.NewDatabase(cnf.DATABASE_ARGS)
	if err := db.Connect(); err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			api.SendBeseResponse(w, false, err)()
		}
	}()
	profile, err := GetProfileByAID(db, AID.(int))
	if err != nil {
		panic(err)
	}
	db.BeginTransaction()
	newQB := qb.NewSyncQB(db.SyncQ()).Delete(cnf.DBT_AUTH).Where(qb.Compare("id", qb.EQUAL, profile.UserId))
	newQB.Merge()
	_, err = newQB.Exec()
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	if err := deleteAvatar(profile); err != nil {
		if err := db.RollBackTransaction(); err != nil {
			return api.SendBeseResponse(w, false, err)
		}
		return api.SendBeseResponse(w, false, err)
	}
	if err := db.CommitTransaction(); err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	return api.SendBeseResponse(w, true, nil)
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
