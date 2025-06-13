package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
)

type UpdateLogGroupForm struct {
	Id          int    `form:"Id" empty:"-err"`
	ProjectId   int    `form:"ProjectId" empty:"-err"`
	Name        string `form:"Name" empty:"-err"`
	Description string `form:"Description"`
}

func UpdateLogGroup(w http.ResponseWriter, r *http.Request, m interfaces.IManager) error {
	UID, ok := m.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}
	var updateLogGroupForm UpdateLogGroupForm
	if err := mapper.FillStructFromForm(frm, &updateLogGroupForm); err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}
	hasPermissions, err := changeProjectPermissions(updateLogGroupForm.ProjectId, UID)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if !hasPermissions {
		return api.NewClientError(http.StatusForbidden, "no access to update the log group")
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
