package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type ProjectMessage struct {
	rest.ImplementDTOMessage
	UserId      string
	Name        string
	Description string
}
