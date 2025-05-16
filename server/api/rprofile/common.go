package rprofile

import (
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
)

func GetProfileByAID(dbRes interfaces.IReadDatabase, AID int) (*ProfileDBView, error) {
	newQB := qb.NewSyncQB(dbRes.SyncQ()).SelectFrom("*", cnf.DBT_PROFILE).Where(qb.Compare("user_id", qb.EQUAL, AID))
	newQB.Merge()
	res, err := newQB.Query()
	if err != nil {
		return nil, err
	}
	profileDBView := make([]ProfileDBView, len(res))
	err = mapper.FillStructSliceFromDb(&profileDBView, &res)
	if err != nil {
		return nil, err
	}
	return &profileDBView[0], nil
}
