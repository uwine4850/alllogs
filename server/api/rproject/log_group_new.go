package rproject

import (
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
	ProjectId   string `form:"ProjectId" empty:"-err"`
	Name        string `form:"Name" empty:"-err"`
	Description string `form:"Description" empty:"-err"`
}

func NewLogGroup(w http.ResponseWriter, r *http.Request, manager interfaces.Manager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	var logGroupForm LogGroupForm
	if err := mapper.FillStructFromForm(frm, &logGroupForm); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	// Database
	projectID, err := strconv.Atoi(logGroupForm.ProjectId)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	isProjectAuthor, err := IsProjectAuthor(UID.(int), projectID, cnf.DatabaseReader)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if !isProjectAuthor {
		return api.NewClientError(http.StatusBadRequest, "permission dained")
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Insert(cnf.DBT_PROJECT_LOG_GROUP,
		map[string]any{
			"project_id": logGroupForm.ProjectId, "name": logGroupForm.Name, "description": logGroupForm.Description,
		})
	newQB.Merge()
	_, err = newQB.Exec()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}
