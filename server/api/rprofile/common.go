package rprofile

import (
	"errors"

	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
)

func GetProfileByUID(dbRes interfaces.IReadDatabase, UID int) (*ProfileDBView, error) {
	newQB := qb.NewSyncQB(dbRes.SyncQ()).SelectFrom("*", cnf.DBT_PROFILE).Where(qb.Compare("user_id", qb.EQUAL, UID))
	newQB.Merge()
	res, err := newQB.Query()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("user not found")
	}
	var profileDBView ProfileDBView
	if err := mapper.FillStructFromDb(&profileDBView, &res[0]); err != nil {
		return nil, err
	}
	return &profileDBView, nil
}
