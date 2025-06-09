package rauth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/interfaces"
)

func Logout(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	logoutForm := mydto.LogoutMessage{}
	if err := json.NewDecoder(r.Body).Decode(&logoutForm); err != nil {
		SendLoginResponse(w, "", 0, err.Error())
		return nil
	}

	_aid, ok := manager.OneTimeData().GetUserContext("AID")
	if !ok {
		api.SendBeseResponse(w, false, errors.New("account logout error"))
		return nil
	}
	AID := _aid.(int)
	if AID != logoutForm.UID {
		api.SendBeseResponse(w, false, errors.New("account logout error"))
		return nil
	} else {
		api.SendBeseResponse(w, true, nil)
		return nil
	}
}
