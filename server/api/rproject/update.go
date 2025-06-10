package rproject

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/config"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Update(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		api.SendBeseResponse(w, false, errors.New("user ID not found"))
		return nil
	}
	updateForm := mydto.ProjectMessage{}
	if err := json.NewDecoder(r.Body).Decode(&updateForm); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}

	dbName := config.LoadedConfig().Default.Database.MainConnectionPoolName
	connection, err := manager.Database().ConnectionPool(dbName)
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	newQB := qb.NewSyncQB(connection.SyncQ())
	newQB.Update(cnf.DBT_PROJECT, map[string]any{
		"name":        updateForm.Name,
		"description": updateForm.Description,
	}).Where(
		qb.Compare("id", qb.EQUAL, updateForm.Id), qb.AND,
		qb.Compare("user_id", qb.EQUAL, UID),
	)
	newQB.Merge()
	res, err := newQB.Exec()
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	if res["rowsAffected"].(int64) == 0 {
		api.SendBeseResponse(w, false, errors.New("error during project update"))
		return nil
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}
