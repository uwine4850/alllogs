package projectperm

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func ProjectPermission(projectId any, m interfaces.Manager, errorText string) error {
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

func EditLogGroupPermission(logGroupId any, UID any) (bool, error) {
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Select(qb.Exists(qb.SQ(
		false,
		qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).SelectFrom("*", cnf.DBT_PROJECT).
			InnerJoin(
				cnf.DBT_PROJECT_LOG_GROUP,
				qb.NoArgsCompare(cnf.DBT_PROJECT+".id", qb.EQUAL, cnf.DBT_PROJECT_LOG_GROUP+".project_id"),
			).
			Where(
				qb.Compare(cnf.DBT_PROJECT_LOG_GROUP+".id", qb.EQUAL, logGroupId), qb.AND,
				qb.Compare(cnf.DBT_PROJECT+".user_id", qb.EQUAL, UID),
			),
	))).As("ok").Merge()
	res, err := newQB.Query()
	if err != nil {
		return false, err
	}
	if len(res) > 0 {
		ok, err := dbutils.ParseInt(res[0]["ok"])
		if err != nil {
			return false, err
		}
		return ok != 0, nil
	} else {
		return false, nil
	}
}
