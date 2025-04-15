package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type BaseResponseMessage struct {
	rest.ImplementDTOMessage
	Ok    bool
	Error string
}

func NewBaseResponse(ok bool, error string) *BaseResponseMessage {
	return &BaseResponseMessage{
		Ok:    ok,
		Error: error,
	}
}
