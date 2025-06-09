package routes

import (
	"github.com/uwine4850/alllogs/api/notifications"
	"github.com/uwine4850/alllogs/api/rauth"
	"github.com/uwine4850/alllogs/api/rprofile"
	"github.com/uwine4850/alllogs/api/rproject"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/object"
)

func Get(database object.IViewDatabase) []map[string]map[string]router.Handler {
	return []map[string]map[string]router.Handler{
		{
			"POST": {"/register": rauth.Register()},
		},
		{
			"POST": {"/login": rauth.Login()},
		},
		{
			"GET": {"/profile/:user_id": rprofile.JsonProfileObjectView(database)},
		},
		{
			"POST": {"/gen-token": rprofile.GenerateToken},
		},
		{
			"DELETE": {"/del-token": rprofile.DeleteToken},
		},
		{
			"GET": {"/profile/update/1": rprofile.JsonProfileObjectView(database)},
		},
		{
			"PUT": {"/profile/update": rprofile.Update},
		},
		{
			"DELETE": {"/profile/del": rprofile.Delete},
		},
		{
			"GET": {"/notifications": notifications.Notification},
		},
		{
			"POST": {"/new-project": rproject.NewProject},
		},
		{
			"GET": {"/project/:id": rproject.ProjectObjectView(database)},
		},
		{
			"POST": {"/new-log-group": rproject.NewLogGroup},
		},
		{
			"GET": {"/project-detail/:projID/log-group/:logID": rproject.LogGroupObjectView(database)},
		},
		{
			"GET": {"/all-projects": rproject.ProjectsAllView(database)},
		},
		{
			"GET": {"/all-log-groups/:project_id": rproject.LogGroupsAllView(database)},
		},
		{
			"POST": {"/logout": rauth.Logout},
		},
	}
}
