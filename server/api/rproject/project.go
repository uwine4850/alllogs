package rproject

import (
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/form/formmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type ProjectForm struct {
	Name        []string `name:"Name"`
	Description []string `name:"Description"`
}

func NewProject(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		return api.SendBeseResponse(w, false, errors.New("User ID not found."))
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	var projectForm ProjectForm
	mapper := formmapper.NewMapper(frm, typeopr.Ptr{}.New(&projectForm), []string{})
	if err := mapper.Fill(); err != nil {
		return api.SendBeseResponse(w, false, err)
	}

	// Database
	db := database.NewDatabase(cnf.DATABASE_ARGS)
	if err := db.Connect(); err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			api.SendBeseResponse(w, false, err)()
		}
	}()

	newQB := qb.NewSyncQB(db.SyncQ()).Insert(cnf.DBT_PROJECT,
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
