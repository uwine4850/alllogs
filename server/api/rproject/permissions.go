package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func ProjectPermission(projectId any, m interfaces.IManager, errorText string) error {
	currentUID, ok := m.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "uid not exists")
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Select(qb.Exists(
		qb.SQ(false, qb.NewNoDbQB().
			SelectFrom(1, cnf.DBT_PROJECT).Where(
			qb.Compare("id", qb.EQUAL, projectId), qb.AND,
			qb.Compare("user_id", qb.EQUAL, currentUID),
		),
		),
	)).As("ok").Merge()
	res, err := newQB.Query()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	if res[0]["ok"].(int64) != 1 {
		return api.NewClientError(http.StatusForbidden, errorText)
	}
	return nil
}
