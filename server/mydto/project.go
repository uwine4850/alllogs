package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type ProjectAuthor struct {
	rest.ImplementDTOMessage
	PID      int    `dto:"PID" db:"pid"`
	Username string `dto:"Username" db:"username"`
	Avatar   string `dto:"Avatar" db:"avatar"`
}

type ProjectMessage struct {
	rest.ImplementDTOMessage
	Author      ProjectAuthor `dto:"Author"`
	Id          int           `dto:"Id" db:"id"`
	UserId      int           `dto:"UserId" db:"user_id"`
	Name        string        `dto:"Name" db:"name"`
	Description string        `dto:"Description" db:"description"`
	Error       string        `dto:"Error"`
}

type ProjectLogGroupMessage struct {
	rest.ImplementDTOMessage
	Id          int    `dto:"Id" db:"id"`
	ProjectId   int    `dto:"ProjectId" db:"project_id"`
	Name        string `dto:"Name" db:"name"`
	Description string `dto:"Description" db:"description"`
	Error       string `dto:"Error"`
}
