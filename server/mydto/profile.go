package mydto

import (
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type ProfileMessage struct {
	rest.ImplementDTOMessage
	User        UserMessage `name:"User"`
	Id          string      `name:"Id"`
	UserId      string      `name:"UserId"`
	Description string      `name:"Description"`
	Avatar      string      `name:"Avatar"`
	Token       string      `name:"Token"`
	Error       string
}

type UserMessage struct {
	rest.ImplementDTOMessage
	Id       string
	Username string
}

type GenTokenMessage struct {
	rest.ImplementDTOMessage
	UserId string
}

type TokenResponse struct {
	rest.ImplementDTOMessage
	Token string
	Error string
}

type ProfileUpdateMessage struct {
	rest.ImplementDTOMessage
	PID           string
	Description   string
	Avatar        form.FormFile
	OldAvatarPath string
	DelAvatar     bool
}
