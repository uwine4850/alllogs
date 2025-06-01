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

func NewLogGroup(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	AID, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		api.SendBeseResponse(w, false, errors.New("user ID not found"))
		return nil
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	var logGroupForm LogGroupForm
	if err := mapper.FillStructFromForm(frm, &logGroupForm); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	// Database
	projectID, err := strconv.Atoi(logGroupForm.ProjectId[0])
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	isProjectAuthor, err := IsProjectAuthor(AID.(int), projectID, cnf.DatabaseReader)
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	if !isProjectAuthor {
		api.SendBeseResponse(w, false, errors.New("permission dained"))
		return nil
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Insert(cnf.DBT_PROJECT_LOG_GROUP,
		map[string]any{
			"project_id": logGroupForm.ProjectId[0], "name": logGroupForm.Name[0], "description": logGroupForm.Description[0],
		})
	newQB.Merge()
	_, err = newQB.Exec()
	if err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}
