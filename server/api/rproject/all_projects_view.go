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

type AllProjectsView struct {
	object.AllView
}

func (v *AllProjectsView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	msg := []mydto.ProjectMessage{
		{
			Error: err.Error(),
		},
	}
	mapper.SendSafeJsonDTOMessage(w, mydto.DTO, typeopr.Ptr{}.New(&msg))
}

func (v *AllProjectsView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func ProjectsAllView() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	view := object.JsonAllTemplateView{
		View: &AllProjectsView{
			object.AllView{
				Name:       "projects",
				TableName:  cnf.DBT_PROJECT,
				FillStruct: mydto.ProjectMessage{},
			},
		},
		DTO:     mydto.DTO,
		Message: mydto.ProjectMessage{},
	}
	return view.Call
}
