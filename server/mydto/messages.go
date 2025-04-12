package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

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
