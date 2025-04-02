package mydto

import (
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

var DTO = rest.NewDTO()

var AllowMessages = []rest.AllowMessage{
	{
		Package: "mydto",
		Name:    "Register",
	},
	{
		Package: "mydto",
		Name:    "BaseResponse",
	},
}

var TSGenMessages = map[string]*[]irest.IMessage{
	"../client/src/dto/auth.ts": {
		Register{},
	},
	"../client/src/dto/common.ts": {
		BaseResponse{},
	},
}

func SetUpMessages(dto *rest.DTO) {
	dto.AllowedMessages(AllowMessages)
	dto.Messages(TSGenMessages)
}
