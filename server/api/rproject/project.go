package rproject

import (
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
)

type ProjectForm struct {
	Name        []string `form:"Name" empty:"-err"`
	Description []string `form:"Description" empty:"-err"`
}

func NewProject(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		api.SendBeseResponse(w, false, errors.New("user ID not found"))
		return nil
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	var projectForm ProjectForm
	if err := mapper.FillStructFromForm(frm, &projectForm); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Insert(cnf.DBT_PROJECT,
		map[string]any{
			"user_id": UID, "name": projectForm.Name[0], "description": projectForm.Description[0],
		})
	newQB.Merge()
	_, err := newQB.Exec()
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}

func IsProjectAuthor(UID int, projectId int, dbRead interfaces.IReadDatabase) (bool, error) {
	newQB := qb.NewSyncQB(dbRead.SyncQ())
	return qb.SelectExists(newQB, cnf.DBT_PROJECT,
		qb.Compare("id", qb.EQUAL, projectId), qb.AND,
		qb.Compare("user_id", qb.EQUAL, UID))
}
