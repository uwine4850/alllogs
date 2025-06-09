package rproject

import (
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type ProjectView struct {
	object.ObjView
}

func (v *ProjectView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	msg := mydto.ProjectMessage{}
	msg.Error = err.Error()
	mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(&msg))
}

func (v *ProjectView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func ProjectObjectView(database object.IViewDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	view := object.JsonObjectTemplateView{
		View: &ProjectView{
			object.ObjView{
				Name:       "object",
				TableName:  cnf.DBT_PROJECT,
				Database:   database,
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
		newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
		newQB.SelectFrom("auth.username, profile.avatar, profile.user_id", cnf.DBT_AUTH).
			InnerJoin(cnf.DBT_PROFILE, qb.NoArgsCompare("user_id", qb.EQUAL, "auth.id")).
			Where(qb.Compare("auth.id", qb.EQUAL, msg.UserId))
		newQB.Merge()
		res, err := newQB.Query()
		if err != nil {
			return err
		}

		author := make([]mydto.ProjectAuthor, len(res))
		err = mapper.FillStructSliceFromDb(&author, &res)
		if err != nil {
			return err
		}
		msg.Author = author[0]
		return nil
	})
	return view.Call
}
