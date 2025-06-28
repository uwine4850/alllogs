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
		Name:    "MsgRegister",
	},
	{
		Package: "rauth",
		Name:    "MsgLogin",
	},
	{
		Package: "rauth",
		Name:    "MsgLoginResponse",
	},
	{
		Package: "api",
		Name:    "MsgBaseResponse",
	},
	{
		Package: "rprofile",
		Name:    "MsgProfile",
	},
	{
		Package: "rprofile",
		Name:    "MsgUser",
	},
	{
		Package: "rprofile",
		Name:    "MsgGenToken",
	},
	{
		Package: "rprofile",
		Name:    "MsgTokenResponse",
	},
	{
		Package: "rprofile",
		Name:    "MsgProfileUpdate",
	},
	{
		Package: "notifications",
		Name:    "MsgNotificationInfo",
	},
	{
		Package: "rproject",
		Name:    "MsgProject",
	},
	{
		Package: "rproject",
		Name:    "MsgProjectAuthor",
	},
	{
		Package: "rproject",
		Name:    "MsgProjectLogGroup",
	},
	{
		Package: "api",
		Name:    "MsgClientError",
	},
	{
		Package: "rauth",
		Name:    "MsgLogout",
	},
	{
		Package: "api",
		Name:    "MsgServerError",
	},
	{
		Package: "rproject",
		Name:    "MsgLogItem",
	},
	{
		Package: "rproject",
		Name:    "MsgLogItemPayload",
	},
	{
		Package: "rproject",
		Name:    "MsgLogItemsFilter",
	},
}

var TSGenMessages = map[string][]irest.IMessage{
	"../client/src/dto/common.ts": {
		api.MsgBaseResponse{},
		api.MsgClientError{},
		api.MsgServerError{},
	},
	"../client/src/dto/auth.ts": {
		rauth.MsgRegister{},
		rauth.MsgLogin{},
		rauth.MsgLoginResponse{},
		rauth.MsgLogout{},
	},
	"../client/src/dto/profile.ts": {
		rprofile.MsgProfile{},
		rprofile.MsgUser{},
		rprofile.MsgGenToken{},
		rprofile.MsgTokenResponse{},
		rprofile.MsgProfileUpdate{},
	},
	"../client/src/dto/notifications.ts": {
		notifications.MsgNotificationInfo{},
	},
	"../client/src/dto/project.ts": {
		rproject.MsgProject{},
		rproject.MsgProjectAuthor{},
		rproject.MsgProjectLogGroup{},
		rproject.MsgLogItem{},
		rproject.MsgLogItemPayload{},
		rproject.MsgLogItemsFilter{},
	},
}

func SetUpMessages(dto *rest.DTO) {
	dto.AllowedMessages(AllowMessages)
	dto.Messages(TSGenMessages)
}
