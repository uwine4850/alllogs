package rprofile

import (
	"fmt"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
)

func GetProfileByAID(db *database.Database, AID string) {
	// profileView := DBRowView{}
	newQB := qb.NewSyncQB(db.SyncQ()).SelectFrom("*", cnf.DBT_PROFILE).Where(qb.Compare("user_id", qb.EQUAL, AID))
	newQB.Merge()
	q, err := newQB.Query()
	if err != nil {
		panic(err)
	}
	fmt.Println(q)
}
