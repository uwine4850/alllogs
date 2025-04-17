package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

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
