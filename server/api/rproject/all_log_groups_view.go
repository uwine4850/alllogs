package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/object"
)

// TODO: add permission

type AllLogGroupsView struct {
	object.AllView
}

func (v *AllLogGroupsView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *AllLogGroupsView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func (v *AllLogGroupsView) Permissions(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (bool, func()) {
	slugProjectId, ok := manager.OneTimeData().GetSlugParams(v.Slug)
	if !ok {
		return false, func() {
			api.SendServerError(w, http.StatusInternalServerError, "no slug")
		}
	}
	if err := ProjectPermission(slugProjectId, manager, "no permission to view log groups"); err != nil {
		return false, func() { api.SendClientError(w, http.StatusForbidden, err.Error()) }
	}
	return true, func() {}
}

func LogGroupsAllView(database object.IViewDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	view := object.JsonAllTemplateView{
		View: &AllLogGroupsView{
			object.AllView{
				Name:       "groups",
				TableName:  cnf.DBT_PROJECT_LOG_GROUP,
				Database:   database,
				Slug:       "project_id",
				FillStruct: MsgProjectLogGroup{},
			},
		},
		DTO:     cnf.DTO,
		Message: MsgProjectLogGroup{},
	}
	return view.Call
}
