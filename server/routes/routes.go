package routes

import (
	"github.com/uwine4850/alllogs/api/notifications"
	"github.com/uwine4850/alllogs/api/rauth"
	"github.com/uwine4850/alllogs/api/rprofile"
	"github.com/uwine4850/alllogs/api/rproject"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/object"
)

func Get(database interfaces.IDatabase) []map[string]map[string]router.Handler {
	objectViewDb := object.NewViewMysqlDatabase(database)
	return []map[string]map[string]router.Handler{
		{
			"POST": {"/register": rauth.Register()},
		},
		{
			"POST": {"/login": rauth.Login()},
		},
		{
			"GET": {"/profile/:user_id": rprofile.JsonProfileObjectView(objectViewDb)},
		},
		{
			"POST": {"/gen-token": rprofile.GenerateToken},
		},
		{
			"DELETE": {"/del-token": rprofile.DeleteToken},
		},
		{
			"PATCH": {"/profile/update": rprofile.Update},
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
			"GET": {"/project/:id": rproject.ProjectObjectView(objectViewDb)},
		},
		{
			"POST": {"/new-log-group": rproject.NewLogGroup},
		},
		{
			"GET": {"/project-detail/:projID/log-group/:logID": rproject.LogGroupObjectView(objectViewDb)},
		},
		{
			"GET": {"/all-projects/:user_id": rproject.ProjectsAllView(objectViewDb)},
		},
		{
			"GET": {"/all-log-groups/:project_id": rproject.LogGroupsAllView(objectViewDb)},
		},
		{
			"POST": {"/logout": rauth.Logout},
		},
		{
			"PATCH": {"/project": rproject.Update},
		},
		{
			"DELETE": {"/project/:id": rproject.Delete},
		},
		{
			"PATCH": {"/log-group": rproject.UpdateLogGroup},
		},
		{
			"DELETE": {"/project/:projId/log-group/:logId": rproject.LogGroupDelete},
		},
		{
			"GET": {"/logitem": rproject.LogClientSocket},
		},
		{
			"GET": {"/log-items/:logGroupId/:start/:count": rproject.LogItemsObjectView(database)},
		},
	}
}
