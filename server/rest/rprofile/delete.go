package rprofile

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/rest"
	"github.com/uwine4850/foozy/pkg/database"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Delete(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		return rest.SendBeseResponse(w, false, errors.New("User ID not found."))
	}
	db := database.NewDatabase(cnf.DATABASE_ARGS)
	if err := db.Connect(); err != nil {
		return rest.SendBeseResponse(w, false, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			rest.SendBeseResponse(w, false, err)()
		}
	}()
	profile, err := GetProfileByAID(db, AID.(string))
	if err != nil {
		panic(err)
	}
	db.BeginTransaction()
	newQB := qb.NewSyncQB(db.SyncQ()).Delete(cnf.DBT_AUTH).Where(qb.Compare("id", qb.EQUAL, profile.UserId))
	newQB.Merge()
	_, err = newQB.Exec()
	if err != nil {
		return rest.SendBeseResponse(w, false, err)
	}
	if err := deleteAvatar(profile); err != nil {
		if err := db.RollBackTransaction(); err != nil {
			return rest.SendBeseResponse(w, false, err)
		}
		return rest.SendBeseResponse(w, false, err)
	}
	if err := db.CommitTransaction(); err != nil {
		return rest.SendBeseResponse(w, false, err)
	}
	return rest.SendBeseResponse(w, true, nil)
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
