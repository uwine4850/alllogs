package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type AllLogGroupsView struct {
	object.AllView
}

func (v *AllLogGroupsView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	msg := []mydto.ProjectLogGroupMessage{
		{
			Error: err.Error(),
		},
	}
	mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(&msg))
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
				FillStruct: mydto.ProjectLogGroupMessage{},
			},
		},
		DTO:     mydto.DTO,
		Message: mydto.ProjectLogGroupMessage{},
	}
	return view.Call
}
