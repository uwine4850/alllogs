package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/permissions/projectperm"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func ClearLogs(w http.ResponseWriter, r *http.Request, m interfaces.IManager) error {
	UID, ok := m.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}
	logGroupIdSlug, ok := m.OneTimeData().GetSlugParams("logGroupID")
	if !ok {
		return api.NewClientError(http.StatusBadRequest, "log group id not found")
	}

	// Check permission.
	hasPermission, err := projectperm.EditLogGroupPermission(logGroupIdSlug, UID)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if !hasPermission {
		return api.NewClientError(http.StatusForbidden, "no access to clear logs")
	}

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Delete(cnf.DBT_LOG_ITEM).Where(
		qb.Compare("id", qb.EQUAL, logGroupIdSlug),
	)
	_, err = newQB.Exec()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}
