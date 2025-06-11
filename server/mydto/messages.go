package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type BaseResponseMessage struct {
	rest.ImplementDTOMessage
	TypBaseResponseMessage rest.TypeId `dto:"-typeid"`
	Ok                     bool        `dto:"Ok"`
	Error                  string      `dto:"Error"`
}

func NewBaseResponse(ok bool, error string) *BaseResponseMessage {
	return &BaseResponseMessage{
		Ok:    ok,
		Error: error,
	}
}

// 400 - 499
type ClientErrorMessage struct {
	rest.ImplementDTOMessage
	TypClientErrorMessage rest.TypeId `dto:"-typeid"`
	Code                  int         `dto:"Code"`
	Text                  string      `dto:"Text"`
}

func NewClientErrorMessage(code int, text string) *ClientErrorMessage {
	return &ClientErrorMessage{
		Code: code,
		Text: text,
	}
}

// 500 - 511
type ServerErrorMessage struct {
	rest.ImplementDTOMessage
	TypServerErrorMessage rest.TypeId `dto:"-typeid"`
	Code                  int         `dto:"Code"`
	Text                  string      `dto:"Text"`
}

func NewServerErrorMessage(code int, text string) *ServerErrorMessage {
	return &ServerErrorMessage{
		Code: code,
		Text: text,
	}
}
