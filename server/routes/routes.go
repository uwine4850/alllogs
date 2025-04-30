package routes

import (
	"github.com/uwine4850/alllogs/api/notifications"
	"github.com/uwine4850/alllogs/api/rauth"
	"github.com/uwine4850/alllogs/api/rprofile"
	"github.com/uwine4850/foozy/pkg/router"
)

func Get() []map[string]map[string]router.Handler {
	return []map[string]map[string]router.Handler{
		{
			router.POST: {"/register": rauth.Register()},
		},
		{
			router.POST: {"/login": rauth.Login()},
		},
		{
			router.GET: {"/isauth": rauth.IsAuth()},
		},
		{
			router.GET: {"/profile/<id>": rprofile.JsonProfileObjectView()},
		},
		{
			router.POST: {"/gen-token": rprofile.GenerateToken},
		},
		{
			router.DELETE: {"/del-token": rprofile.DeleteToken},
		},
		{
			router.GET: {"/profile/update/1": rprofile.JsonProfileObjectView()},
		},
		{
			router.PUT: {"/profile/update": rprofile.Update},
		},
		{
			router.DELETE: {"/profile/del": rprofile.Delete},
		},
		{
			router.WS: {"/notifications": notifications.Notification},
		},
	}
}
