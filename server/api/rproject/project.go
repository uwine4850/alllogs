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

func NewProject(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		return api.SendBeseResponse(w, false, errors.New("user ID not found"))
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	var projectForm ProjectForm
	if err := mapper.FillStructFromForm(frm, &projectForm); err != nil {
		return api.SendBeseResponse(w, false, err)
	}

	// Database
	// db := database.NewDatabase(cnf.DATABASE_ARGS)
	// if err := db.Connect(); err != nil {
	// 	return api.SendBeseResponse(w, false, err)
	// }
	// defer func() {
	// 	if err := db.Close(); err != nil {
	// 		api.SendBeseResponse(w, false, err)()
	// 	}
	// }()

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Insert(cnf.DBT_PROJECT,
		map[string]any{
			"user_id": AID, "name": projectForm.Name[0], "description": projectForm.Description[0],
		})
	newQB.Merge()
	_, err := newQB.Exec()
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	return api.SendBeseResponse(w, true, nil)
}

func IsProjectAuthor(AID int, projectId int, dbRead interfaces.IReadDatabase) (bool, error) {
	newQB := qb.NewSyncQB(dbRead.SyncQ())
	return qb.SelectExists(newQB, cnf.DBT_PROJECT,
		qb.Compare("id", qb.EQUAL, projectId), qb.AND,
		qb.Compare("user_id", qb.EQUAL, AID))
}
