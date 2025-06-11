package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/config"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
)

type updateForm struct {
	Id          int    `form:"Id"`
	Name        string `form:"Name"`
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

	dbName := config.LoadedConfig().Default.Database.MainConnectionPoolName
	connection, err := manager.Database().ConnectionPool(dbName)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
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
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if res["rowsAffected"].(int64) == 0 {
		return api.NewServerError(http.StatusInternalServerError, "error during project update")
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}
