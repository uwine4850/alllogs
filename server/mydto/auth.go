package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type Register struct {
	rest.ImplementDTOMessage
	Username       string
	Password       string
	RepeatPassword string
}

type Login struct {
	rest.ImplementDTOMessage
	Username string
	Password string
}

type LoginResponse struct {
	rest.ImplementDTOMessage
	JWT   string
	Error string
}
