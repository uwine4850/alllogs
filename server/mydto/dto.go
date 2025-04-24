package mydto

import (
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

var DTO = rest.NewDTO()

var AllowMessages = []rest.AllowMessage{
	{
		Package: "mydto",
		Name:    "RegisterMessage",
	},
	{
		Package: "mydto",
		Name:    "LoginMessage",
	},
	{
		Package: "mydto",
		Name:    "LoginResponseMessage",
	},
	{
		Package: "mydto",
		Name:    "BaseResponseMessage",
	},
	{
		Package: "mydto",
		Name:    "ProfileMessage",
	},
	{
		Package: "mydto",
		Name:    "UserMessage",
	},
	{
		Package: "mydto",
		Name:    "GenTokenMessage",
	},
	{
		Package: "mydto",
		Name:    "TokenResponse",
	},
	{
		Package: "mydto",
		Name:    "AuthorizationMessage",
	},
}

var TSGenMessages = map[string]*[]irest.IMessage{
	"../client/src/dto/common.ts": {
		BaseResponseMessage{},
	},
	"../client/src/dto/auth.ts": {
		RegisterMessage{},
		LoginMessage{},
		LoginResponseMessage{},
		AuthorizationMessage{},
	},
	"../client/src/dto/profile.ts": {
		ProfileMessage{},
		UserMessage{},
		GenTokenMessage{},
		TokenResponse{},
	},
}

func SetUpMessages(dto *rest.DTO) {
	dto.AllowedMessages(AllowMessages)
	dto.Messages(TSGenMessages)
}
