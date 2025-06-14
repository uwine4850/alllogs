package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/object"
)

type AllLogGroupsView struct {
	object.AllView
}

func (v *AllLogGroupsView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *AllLogGroupsView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func LogGroupsAllView(database object.IViewDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	view := object.JsonAllTemplateView{
		View: &AllLogGroupsView{
			object.AllView{
				Name:       "groups",
				TableName:  cnf.DBT_PROJECT_LOG_GROUP,
				Database:   database,
				Slug:       "project_id",
				FillStruct: ProjectLogGroupMessage{},
			},
		},
		DTO:     cnf.DTO,
		Message: ProjectLogGroupMessage{},
	}
	return view.Call
}
