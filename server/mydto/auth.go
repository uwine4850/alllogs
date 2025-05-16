package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type RegisterMessage struct {
	rest.ImplementDTOMessage
	Username       string `dto:"Username"`
	Password       string `dto:"Password"`
	RepeatPassword string `dto:"RepeatPassword"`
}

type LoginMessage struct {
	rest.ImplementDTOMessage
	Username string `dto:"Username"`
	Password string `dto:"Password"`
}

type LoginResponseMessage struct {
	rest.ImplementDTOMessage
	JWT   string `dto:"JWT"`
	UID   int    `dto:"UID"`
	Error string `dto:"Error"`
}
