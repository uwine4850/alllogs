package rproject

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
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
	if err := mapper.FillStructFromForm(frm, &logGroupForm); err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	// Database
	projectID, err := strconv.Atoi(logGroupForm.ProjectId[0])
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	isProjectAuthor, err := IsProjectAuthor(AID.(int), projectID, cnf.DatabaseReader)
	if err != nil {
		return api.SendBeseResponse(w, false, err)
	}
	if !isProjectAuthor {
		return api.SendBeseResponse(w, false, errors.New("permission dained"))
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Insert(cnf.DBT_PROJECT_LOG_GROUP,
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
