package mydto

import (
	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/notifications"
	"github.com/uwine4850/alllogs/api/rauth"
	"github.com/uwine4850/alllogs/api/rprofile"
	"github.com/uwine4850/alllogs/api/rproject"
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

var AllowMessages = []rest.AllowMessage{
	{
		Package: "rauth",
		Name:    "RegisterMessage",
	},
	{
		Package: "rauth",
		Name:    "LoginMessage",
	},
	{
		Package: "rauth",
		Name:    "LoginResponseMessage",
	},
	{
		Package: "api",
		Name:    "BaseResponseMessage",
	},
	{
		Package: "rprofile",
		Name:    "ProfileMessage",
	},
	{
		Package: "rprofile",
		Name:    "UserMessage",
	},
	{
		Package: "rprofile",
		Name:    "GenTokenMessage",
	},
	{
		Package: "rprofile",
		Name:    "TokenResponse",
	},
	{
		Package: "rprofile",
		Name:    "ProfileUpdateMessage",
	},
	{
		Package: "notifications",
		Name:    "NotificationInfoMessage",
	},
	{
		Package: "rproject",
		Name:    "ProjectMessage",
	},
	{
		Package: "rproject",
		Name:    "ProjectAuthor",
	},
	{
		Package: "rproject",
		Name:    "ProjectLogGroupMessage",
	},
	{
		Package: "api",
		Name:    "ClientErrorMessage",
	},
	{
		Package: "rauth",
		Name:    "LogoutMessage",
	},
	{
		Package: "api",
		Name:    "ServerErrorMessage",
	},
	{
		Package: "rproject",
		Name:    "LogItemMessage",
	},
	{
		Package: "rproject",
		Name:    "LogItemPayload",
	},
	{
		Package: "rproject",
		Name:    "LogItemsFilterMessage",
	},
}

var TSGenMessages = map[string][]irest.IMessage{
	"../client/src/dto/common.ts": {
		api.BaseResponseMessage{},
		api.ClientErrorMessage{},
		api.ServerErrorMessage{},
	},
	"../client/src/dto/auth.ts": {
		rauth.RegisterMessage{},
		rauth.LoginMessage{},
		rauth.LoginResponseMessage{},
		rauth.LogoutMessage{},
	},
	"../client/src/dto/profile.ts": {
		rprofile.ProfileMessage{},
		rprofile.UserMessage{},
		rprofile.GenTokenMessage{},
		rprofile.TokenResponse{},
		rprofile.ProfileUpdateMessage{},
	},
	"../client/src/dto/notifications.ts": {
		notifications.NotificationInfoMessage{},
	},
	"../client/src/dto/project.ts": {
		rproject.ProjectMessage{},
		rproject.ProjectAuthor{},
		rproject.ProjectLogGroupMessage{},
		rproject.LogItemMessage{},
		rproject.LogItemPayload{},
		rproject.LogItemsFilterMessage{},
	},
}

func SetUpMessages(dto *rest.DTO) {
	dto.AllowedMessages(AllowMessages)
	dto.Messages(TSGenMessages)
}
