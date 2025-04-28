package rprofile

import (
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

func GetProfileByAID(db *database.Database, AID string) (*ProfileDBView, error) {
	newQB := qb.NewSyncQB(db.SyncQ()).SelectFrom("*", cnf.DBT_PROFILE).Where(qb.Compare("user_id", qb.EQUAL, AID))
	newQB.Merge()
	q, err := newQB.Query()
	if err != nil {
		return nil, err
	}
	profileDBView := []ProfileDBView{}
	mapper := dbmapper.NewMapper(q, typeopr.Ptr{}.New(&profileDBView))
	if err := mapper.Fill(); err != nil {
		return nil, err
	}
	return &profileDBView[0], nil
}
