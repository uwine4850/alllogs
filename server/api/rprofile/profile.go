package rprofile

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/object"
)

type ProfileDBView struct {
	UserId      int    `db:"user_id"`
	Description string `db:"description"`
	Avatar      string `db:"avatar"`
	Token       string `db:"token"`
}

type JsonProfileObject struct {
	object.ObjView
}

func (v *JsonProfileObject) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *JsonProfileObject) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func JsonProfileObjectView(database object.IViewDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	view := object.JsonObjectTemplateView{
		View: &JsonProfileObject{
			object.ObjView{
				Name:       "object",
				TableName:  cnf.DBT_PROFILE,
				Database:   database,
				FillStruct: ProfileDBView{},
				Slug:       "user_id",
			},
		},
		DTO:     mydto.DTO,
		Message: mydto.ProfileMessage{},
	}
	view.OnMessageFilled(func(message any, manager interfaces.IManager) error {
		profileMessage := message.(*mydto.ProfileMessage)
		querybuild := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).SelectFrom("id, username", cnf.DBT_AUTH).
			Where(qb.Compare("id", qb.EQUAL, profileMessage.UserId))
		querybuild.Merge()
		userData, err := querybuild.Query()
		if err != nil {
			return err
		}

		authUser := make([]auth.User, len(userData))
		err = mapper.FillStructSliceFromDb(&authUser, &userData)
		if err != nil {
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
