package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type RegisterMessage struct {
	rest.ImplementDTOMessage
	TypRegisterMessage rest.TypeId `dto:"-typeid"`
	Username           string      `dto:"Username"`
	Password           string      `dto:"Password"`
	RepeatPassword     string      `dto:"RepeatPassword"`
}

type LoginMessage struct {
	rest.ImplementDTOMessage
	TypLoginMessage rest.TypeId `dto:"-typeid"`
	Username        string      `dto:"Username"`
	Password        string      `dto:"Password"`
}

type LoginResponseMessage struct {
	rest.ImplementDTOMessage
	TypLoginResponseMessage rest.TypeId `dto:"-typeid"`
	JWT                     string      `dto:"JWT"`
	UID                     int         `dto:"UID"`
	Error                   string      `dto:"Error"`
}

type LogoutMessage struct {
	rest.ImplementDTOMessage
	TypLogoutMessage rest.TypeId `dto:"-typeid"`
	UID              int         `dto:"UID"`
}
