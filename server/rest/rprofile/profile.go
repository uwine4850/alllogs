package rprofile

import (
	"net/http"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/database/dbmapper"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/namelib"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/router/rest/restmapper"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type DBRowView struct {
	Id          string `name:"id"`
	UserId      string `name:"user_id"`
	Description string `name:"description"`
	Avatar      string `name:"avatar"`
	Token       string `name:"token"`
}

type JsonProfileObject struct {
	object.ObjView
}

func (v *JsonProfileObject) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	msg := mydto.ProfileMessage{}
	msg.Error = err.Error()
	restmapper.SendSafeJsonMessage(w, mydto.DTO, typeopr.Ptr{}.New(&msg))
}

func (v *JsonProfileObject) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func JsonProfileObjectView() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DATABASE_ARGS)
	view := object.JsonObjectTemplateView{
		View: &JsonProfileObject{
			object.ObjView{
				Name:       "object",
				DB:         db,
				TableName:  cnf.DBT_PROFILE,
				FillStruct: DBRowView{},
				Slug:       "id",
			},
		},
		DTO:     mydto.DTO,
		Message: mydto.ProfileMessage{},
	}
	view.OnMessageFilled(func(message any, manager interfaces.IManager) error {
		profileMessage := message.(*mydto.ProfileMessage)
		_db, _ := manager.OneTimeData().GetUserContext(namelib.OBJECT.OBJECT_DB)
		db := _db.(*database.Database)
		querybuild := qb.NewSyncQB(db.SyncQ()).SelectFrom("id, username", cnf.DBT_AUTH).Where(qb.Compare("id", qb.EQUAL, profileMessage.UserId))
		querybuild.Merge()
		userData, err := querybuild.Query()
		if err != nil {
			return err
		}
		var authUser []auth.User
		mapper := dbmapper.NewMapper(userData, typeopr.Ptr{}.New(&authUser))
		if err := mapper.Fill(); err != nil {
			return err
		}
		if len(authUser) == 1 {
			userMessage := mydto.UserMessage{
				Id:       authUser[0].Id,
				Username: authUser[0].Username,
			}
			profileMessage.User = userMessage
		}
		return nil
	})
	return view.Call
}
