package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type ProjectAuthor struct {
	rest.ImplementDTOMessage
	PID      int    `db:"pid"`
	Username string `db:"username"`
	Avatar   string `db:"avatar"`
}

type ProjectMessage struct {
	rest.ImplementDTOMessage
	Author      ProjectAuthor
	UserId      int    `db:"user_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Error       string
}
