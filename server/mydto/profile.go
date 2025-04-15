package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type ProfileMessage struct {
	rest.ImplementDTOMessage
	User        UserMessage `json:"User"`
	Id          string      `json:"Id"`
	UserId      string      `json:"UserId"`
	Description string      `json:"Description"`
	Avatar      string      `json:"Avatar"`
	Token       string      `json:"Token"`
	Error       string
}

type UserMessage struct {
	rest.ImplementDTOMessage
	Id       string
	Username string
}
