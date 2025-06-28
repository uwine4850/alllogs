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
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	updateForm := updateForm{}
	if err := mapper.FillStructFromForm(frm, &updateForm); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	hasPermission, err := changeProjectPermissions(updateForm.Id, UID)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if !hasPermission {
		return api.NewClientError(http.StatusForbidden, "no access to update the project")
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
