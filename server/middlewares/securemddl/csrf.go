package securemddl

import (
	"net/http"
	"slices"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/secure"
)

var skipUrl = []string{"/set-csrf", "/notifications", "/logitem"}

func ValidateCSRFToken(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	if slices.Contains(skipUrl, r.URL.Path) {
		return nil
	}
	if err := secure.ValidateHeaderCSRFToken(r, "X-CSRF-TOKEN"); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
