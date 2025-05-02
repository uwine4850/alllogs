package rproject

import (
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/router/rest/restmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type ProjectView struct {
	object.ObjView
}

func (v *ProjectView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	msg := mydto.ProjectMessage{}
	msg.Error = err.Error()
	restmapper.SendSafeJsonMessage(w, mydto.DTO, typeopr.Ptr{}.New(&msg))
}

func (v *ProjectView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func ProjectObjectView() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DATABASE_ARGS)
	view := object.JsonObjectTemplateView{
		View: &ProjectView{
			object.ObjView{
				Name:       "object",
				DB:         db,
				TableName:  cnf.DBT_PROJECT,
				FillStruct: mydto.ProjectMessage{},
				Slug:       "id",
			},
		},
		DTO:     mydto.DTO,
		Message: mydto.ProjectMessage{},
	}
	view.OnMessageFilled(func(message any, manager interfaces.IManager) error {
		msg, ok := message.(*mydto.ProjectMessage)
		if !ok {
			return errors.New(("error converting a filled message"))
		}
		_db, ok := manager.OneTimeData().GetUserContext(namelib.OBJECT.OBJECT_DB)
		if !ok {
			return errors.New(("error retrieving database from object"))
		}
		db := _db.(*database.Database)
		newQB := qb.NewSyncQB(db.SyncQ())
		newQB.SelectFrom("auth.username, profile.avatar, profile.id as pid", cnf.DBT_AUTH).
			InnerJoin(cnf.DBT_PROFILE, qb.NoArgsCompare("user_id", qb.EQUAL, "auth.id")).
			Where(qb.Compare("auth.id", qb.EQUAL, msg.UserId))
		newQB.Merge()
		res, err := newQB.Query()
		if err != nil {
			return err
		}
		var author []mydto.ProjectAuthor
		dbMapper := dbmapper.NewMapper(res, typeopr.Ptr{}.New(&author))
		if err := dbMapper.Fill(); err != nil {
			return err
		}
		msg.Author = author[0]
		return nil
	})
	return view.Call
}
