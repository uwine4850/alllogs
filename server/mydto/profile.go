package mydto

import (
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type ProfileMessage struct {
	rest.ImplementDTOMessage
	User        UserMessage `dto:"User"`
	Id          int         `dto:"Id"`
	UserId      int         `dto:"UserId"`
	Description string      `dto:"Description"`
	Avatar      string      `dto:"Avatar"`
	Token       string      `dto:"Token"`
	Error       string      `dto:"Error"`
}

type UserMessage struct {
	rest.ImplementDTOMessage
	Id       int    `dto:"Id"`
	Username string `dto:"Username"`
}

type GenTokenMessage struct {
	rest.ImplementDTOMessage
	UserId int `dto:"UserId"`
}

type TokenResponse struct {
	rest.ImplementDTOMessage
	Token string `dto:"Token"`
	Error string `dto:"Error"`
}

type ProfileUpdateMessage struct {
	rest.ImplementDTOMessage
	PID           int           `dto:"PID"`
	Description   string        `dto:"Description"`
	Avatar        form.FormFile `dto:"Avatar"`
	OldAvatarPath string        `dto:"OldAvatarPath"`
	DelAvatar     bool          `dto:"DelAvatar"`
}
