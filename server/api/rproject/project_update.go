package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/apiform"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

type updateForm struct {
	Id          int    `form:"Id" empty:"-err"`
	Name        string `form:"Name" empty:"-err"`
	Description string `form:"Description"`
}

func Update(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}
	updateForm := updateForm{}
	if err := apiform.ParseAndFill(r, &updateForm); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if err := ProjectPermission(updateForm.Id, manager, "no access to update the project"); err != nil {
		return err
	}

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
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
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if res["rowsAffected"].(int64) == 0 {
		return api.NewServerError(http.StatusInternalServerError, "There has been no change")
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}
