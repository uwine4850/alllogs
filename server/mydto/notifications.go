package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type NotificationInfoMessage struct {
	rest.ImplementDTOMessage
	PID  string `dto:"PID"`
	Text string `dto:"Text"`
}
