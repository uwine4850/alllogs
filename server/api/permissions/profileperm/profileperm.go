package profileperm

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func ProfilePermission(m interfaces.IManager, UID int, errorText string) error {
	currentUID, ok := m.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "uid not exists")
	}
	if currentUID.(int) != UID {
		return api.NewClientError(http.StatusForbidden, errorText)
	} else {
		return nil
	}
}
