package rproject

import (
	"encoding/json"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/permissions/projectperm"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
)

func ExportJson(w http.ResponseWriter, r *http.Request, m interfaces.IManager) error {
	UID, ok := m.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}
	logGroupIdSlug, ok := m.OneTimeData().GetSlugParams("logGroupID")
	if !ok {
		return api.NewClientError(http.StatusBadRequest, "log group id not found")
	}

	hasPermission, err := projectperm.EditLogGroupPermission(logGroupIdSlug, UID)
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if !hasPermission {
		return api.NewClientError(http.StatusForbidden, "no access to export logs")
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.SelectFrom("*", cnf.DBT_LOG_ITEM).Where(
		qb.Compare("log_group_id", qb.EQUAL, logGroupIdSlug),
	).Merge()
	res, err := newQB.Query()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	logItems := make([]MsgLogItemPayload, len(res))
	if err := mapper.FillStructSliceFromDb(&logItems, &res); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	jsonData, err := json.MarshalIndent(logItems, "", "  ")
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	w.Header().Set("Content-Disposition", "attachment; filename=\"data.json\"")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonData)
	return nil
}
