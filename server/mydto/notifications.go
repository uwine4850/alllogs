package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type NotificationInfoMessage struct {
	rest.ImplementDTOMessage
	TypNotificationInfoMessage rest.TypeId `dto:"-typeid"`
	PID                        string      `dto:"PID"`
	Text                       string      `dto:"Text"`
}
