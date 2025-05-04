package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type RegisterMessage struct {
	rest.ImplementDTOMessage
	Username       string
	Password       string
	RepeatPassword string
}

type LoginMessage struct {
	rest.ImplementDTOMessage
	Username string
	Password string
}

type LoginResponseMessage struct {
	rest.ImplementDTOMessage
	JWT   string
	UID   int
	Error string
}
