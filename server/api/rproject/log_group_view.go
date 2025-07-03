package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/permissions/projectperm"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/database/dbutils"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type MsgProjectLogGroup struct {
	rest.ImplementDTOMessage
	TypProjectLogGroupMessage rest.TypeId `dto:"-typeid"`
	Id                        int         `dto:"Id" db:"id"`
	ProjectId                 int         `dto:"ProjectId" db:"project_id"`
	Name                      string      `dto:"Name" db:"name"`
	Description               string      `dto:"Description" db:"description"`
	Error                     string      `dto:"Error"`
	AuthorToken               string      `dto:"AuthorToken"`
}

type LogGroupView struct {
	object.MultipleObjectView
}

func (v *LogGroupView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *LogGroupView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func (v *LogGroupView) Permissions(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (bool, func()) {
	slugProjectId, ok := manager.OneTimeData().GetSlugParams("projID")
	if !ok {
		return false, func() {
			api.SendServerError(w, http.StatusInternalServerError, "no slug")
		}
	}
	if err := projectperm.ProjectPermission(slugProjectId, manager, "no permission to view log groups"); err != nil {
		return false, func() { api.SendClientError(w, http.StatusForbidden, err.Error()) }
	}
	return true, func() {}
}

func LogGroupObjectView(database object.IViewDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	view := object.JsonMultipleObjectTemplateView{
		View: &LogGroupView{
			object.MultipleObjectView{
				Database: database,
				MultipleObjects: []object.MultipleObject{
					{
						Name:       "project",
						TableName:  cnf.DBT_PROJECT,
						SlugName:   "projID",
						SlugField:  "id",
						FillStruct: MsgProject{},
					},
					{
						Name:       "log",
						TableName:  cnf.DBT_PROJECT_LOG_GROUP,
						SlugName:   "logID",
						SlugField:  "id",
						FillStruct: MsgProjectLogGroup{},
					},
				},
			},
		},
		DTO: cnf.DTO,
		Messages: map[string]irest.IMessage{
			"project": MsgProject{},
			"log":     MsgProjectLogGroup{},
		},
	}
	view.OnMessageFilled(func(message any, manager interfaces.IManager) error {
		logGroupMessage, ok := message.(*MsgProjectLogGroup)
		if !ok {
			return nil
		}

		newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
		newQB.SelectFrom(cnf.DBT_PROFILE+".token", cnf.DBT_PROFILE).
			InnerJoin(cnf.DBT_PROJECT, qb.Compare(cnf.DBT_PROJECT+".id", qb.EQUAL, logGroupMessage.ProjectId)).
			Where(qb.NoArgsCompare(cnf.DBT_PROFILE+".user_id", qb.EQUAL, cnf.DBT_PROJECT+".user_id")).Merge()
		res, err := newQB.Query()
		if err != nil {
			return err
		}
		if len(res) == 0 {
			return object.ErrNoData{}
		}
		token := dbutils.ParseString(res[0]["token"])
		logGroupMessage.AuthorToken = token
		return nil
	})
	return view.Call
}
