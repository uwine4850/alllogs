package rauth

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
)

func IsAuth() router.Handler {
	return func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
		return nil
	}
}
