package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type Register struct {
	rest.ImplementDTOMessage
	Username       string
	Password       string
	RepeatPassword string
}

type BaseResponse struct {
	rest.ImplementDTOMessage
	Ok    bool
	Error string
}

func NewBaseResponse(ok bool, error string) *BaseResponse {
	return &BaseResponse{
		Ok:    ok,
		Error: error,
	}
}
