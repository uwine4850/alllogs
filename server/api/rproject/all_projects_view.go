package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/object"
)

type AllProjectsView struct {
	object.AllView
}

func (v *AllProjectsView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *AllProjectsView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func ProjectsAllView(database object.IViewDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	view := object.JsonAllTemplateView{
		View: &AllProjectsView{
			object.AllView{
				Name:       "projects",
				TableName:  cnf.DBT_PROJECT,
				Database:   database,
				Slug:       "user_id",
				FillStruct: MsgProject{},
			},
		},
		DTO:     cnf.DTO,
		Message: MsgProject{},
	}
	return view.Call
}
