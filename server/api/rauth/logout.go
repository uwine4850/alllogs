package rauth

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type MsgLogout struct {
	rest.ImplementDTOMessage
	TypLogoutMessage rest.TypeId `dto:"-typeid"`
	UID              int         `dto:"UID"`
}

type LogoutForm struct {
	UID int `form:"UID" empty:"-err"`
}

func Logout(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	_aid, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "UID not exists")
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}
	var logoutForm LogoutForm
	if err := mapper.FillStructFromForm(frm, &logoutForm); err != nil {
		return api.NewClientError(http.StatusBadRequest, err.Error())
	}

	AID := _aid.(int)
	if AID != logoutForm.UID {
		return api.NewClientError(http.StatusConflict, "id dont match")
	} else {
		api.SendBeseResponse(w, true, nil)
		return nil
	}
}
