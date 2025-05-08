package rproject

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/form/formmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type LogGroupForm struct {
	ProjectId   []string `form:"ProjectId" empty:"-err"`
	Name        []string `form:"Name" empty:"-err"`
	Description []string `form:"Description" empty:"-err"`
}

func NewLogGroup(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		return api.SendBeseResponse(w, false, errors.New("user ID not found"))
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	var logGroupForm LogGroupForm
	mapper := formmapper.NewMapper(frm, typeopr.Ptr{}.New(&logGroupForm), []string{})
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
	projectID, err := strconv.Atoi(logGroupForm.ProjectId[0])
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	isProjectAuthor, err := IsProjectAuthor(AID.(int), projectID, db)
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	if !isProjectAuthor {
		return api.SendBeseResponse(w, false, errors.New("permission dained"))
	}
	newQB := qb.NewSyncQB(db.SyncQ()).Insert(cnf.DBT_PROJECT_LOG_GROUP,
		map[string]any{
			"project_id": logGroupForm.ProjectId[0], "name": logGroupForm.Name[0], "description": logGroupForm.Description[0],
		})
	newQB.Merge()
	_, err = newQB.Exec()
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	return api.SendBeseResponse(w, true, nil)
}
