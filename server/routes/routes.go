package routes

import (
	"github.com/uwine4850/alllogs/rest/rauth"
	"github.com/uwine4850/foozy/pkg/router"
)

func Get() []map[string]map[string]router.Handler {
	return []map[string]map[string]router.Handler{
		{
			router.POST: {"/register": rauth.Register()},
		},
	}
}
