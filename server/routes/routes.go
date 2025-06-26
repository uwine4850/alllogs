package routes

import (
	"github.com/uwine4850/alllogs/api/notifications"
	"github.com/uwine4850/alllogs/api/rauth"
	"github.com/uwine4850/alllogs/api/rprofile"
	"github.com/uwine4850/alllogs/api/rproject"
	"github.com/uwine4850/alllogs/api/security"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/object"
)

func Get(database interfaces.IDatabase) map[string][]map[string]router.Handler {
	objectViewDb := object.NewViewMysqlDatabase(database)
	return map[string][]map[string]router.Handler{
		router.MethodGET: {
			{"/profile/:user_id": rprofile.JsonProfileObjectView(objectViewDb)},
			{"/notifications": notifications.Notification},
			{"/project/:id": rproject.ProjectObjectView(objectViewDb)},
			{"/project-detail/:projID/log-group/:logID": rproject.LogGroupObjectView(objectViewDb)},
			{"/all-projects/:user_id": rproject.ProjectsAllView(objectViewDb)},
			{"/all-log-groups/:project_id": rproject.LogGroupsAllView(objectViewDb)},
			{"/logitem": rproject.LogClientSocket},
			{"/log-items/:logGroupId/:start/:count": rproject.LogItemsObjectView(database)},
			{"/set-csrf": security.SetCSRFToken},
			{"/logs-export-json/:logGroupID": rproject.ExportJson},
		},
		router.MethodPOST: {
			{"/register": rauth.Register()},
			{"/login": rauth.Login()},
			{"/logout": rauth.Logout},
			{"/gen-token": rprofile.GenerateToken},
			{"/new-project": rproject.NewProject},
			{"/new-log-group": rproject.NewLogGroup},
		},
		router.MethodDELETE: {
			{"/del-token": rprofile.DeleteToken},
			{"/profile/del": rprofile.Delete},
			{"/project/:id": rproject.Delete},
			{"/project/:projId/log-group/:logId": rproject.LogGroupDelete},
		},
		router.MethodPATCH: {
			{"/profile/update": rprofile.Update},
			{"/project": rproject.Update},
			{"/log-group": rproject.UpdateLogGroup},
		},
	}
}
