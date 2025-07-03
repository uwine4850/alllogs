package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/apiform"
	"github.com/uwine4850/alllogs/api/permissions/projectperm"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

type UpdateLogGroupForm struct {
	Id          int    `form:"Id" empty:"-err"`
	ProjectId   int    `form:"ProjectId" empty:"-err"`
	Name        string `form:"Name" empty:"-err"`
	Description string `form:"Description"`
}

func UpdateLogGroup(w http.ResponseWriter, r *http.Request, m interfaces.IManager) error {
	var updateLogGroupForm UpdateLogGroupForm
	if err := apiform.ParseAndFill(r, &updateLogGroupForm); err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}
	if err := projectperm.ProjectPermission(updateLogGroupForm.ProjectId, m, "no access to update the log group"); err != nil {
		return err
	}

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Update(cnf.DBT_PROJECT_LOG_GROUP, map[string]any{
		"name":        updateLogGroupForm.Name,
		"description": updateLogGroupForm.Description,
	}).Where(
		qb.Compare("id", qb.EQUAL, updateLogGroupForm.Id), qb.AND,
		qb.Compare("project_id", qb.EQUAL, updateLogGroupForm.ProjectId),
	).Merge()
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
