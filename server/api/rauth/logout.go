package rauth

import (
	"encoding/json"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Logout(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	logoutForm := mydto.LogoutMessage{}
	if err := json.NewDecoder(r.Body).Decode(&logoutForm); err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}

	_aid, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "AID not exists")
	}
	AID := _aid.(int)
	if AID != logoutForm.UID {
		return api.NewClientError(http.StatusConflict, "id dont match")
	} else {
		api.SendBeseResponse(w, true, nil)
		return nil
	}
}
