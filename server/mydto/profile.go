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
	Error       string
}

type UserMessage struct {
	rest.ImplementDTOMessage
	Id       int
	Username string
}

type GenTokenMessage struct {
	rest.ImplementDTOMessage
	UserId int
}

type TokenResponse struct {
	rest.ImplementDTOMessage
	Token string
	Error string
}

type ProfileUpdateMessage struct {
	rest.ImplementDTOMessage
	PID           int
	Description   string
	Avatar        form.FormFile
	OldAvatarPath string
	DelAvatar     bool
}
