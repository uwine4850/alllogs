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

func Delete(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}

	idSlug, ok := manager.OneTimeData().GetSlugParams("id")
	if !ok {
		return api.NewClientError(http.StatusBadRequest, "project id not found")
	}
	projectId, err := strconv.Atoi(idSlug)
	if err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}
	if err := projectperm.ProjectPermission(projectId, manager, "no access to delete the project"); err != nil {
		return err
	}

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Delete(cnf.DBT_PROJECT).Where(
		qb.Compare("id", qb.EQUAL, projectId), qb.AND,
		qb.Compare("user_id", qb.EQUAL, UID),
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
