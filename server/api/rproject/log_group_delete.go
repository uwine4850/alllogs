package rproject

import (
	"net/http"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/permissions/projectperm"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func LogGroupDelete(w http.ResponseWriter, r *http.Request, manager interfaces.Manager) error {
	projectIdSlug, ok := manager.OneTimeData().GetSlugParams("projId")
	if !ok {
		return api.NewClientError(http.StatusBadRequest, "project id not found")
	}
	logIdSlug, ok := manager.OneTimeData().GetSlugParams("logId")
	if !ok {
		return api.NewClientError(http.StatusBadRequest, "log group id not found")
	}
	projectId, err := strconv.Atoi(projectIdSlug)
	if err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}
	if err := projectperm.ProjectPermission(projectId, manager, "no access to delete the log group"); err != nil {
		return err
	}

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Delete(cnf.DBT_PROJECT_LOG_GROUP).Where(
		qb.Compare("id", qb.EQUAL, logIdSlug),
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
