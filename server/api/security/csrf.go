package security

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/secure"
)

func SetCSRFToken(w http.ResponseWriter, r *http.Request, m interfaces.Manager) error {
	if err := secure.SetCSRFToken(3600, false, w, r, m); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}
