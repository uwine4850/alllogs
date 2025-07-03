package rproject

import (
	"net/http"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/permissions/profileperm"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/object"
)

// TODO: add permission

type AllProjectsView struct {
	object.AllView
}

func (v *AllProjectsView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *AllProjectsView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func (v *AllProjectsView) Permissions(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (bool, func()) {
	slugUserId, ok := manager.OneTimeData().GetSlugParams(v.Slug)
	if !ok {
		return false, func() {
			api.SendServerError(w, http.StatusInternalServerError, "no slug")
		}
	}
	intSlugUserId, err := strconv.Atoi(slugUserId)
	if err != nil {
		return false, func() {
			api.SendServerError(w, http.StatusInternalServerError, err.Error())
		}
	}
	if err := profileperm.ProfilePermission(manager, intSlugUserId, "no permission to view log groups"); err != nil {
		return false, func() {
			api.SendClientError(w, http.StatusForbidden, err.Error())
		}
	}
	return true, func() {}
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
