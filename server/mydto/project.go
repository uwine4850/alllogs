package mydto

import "github.com/uwine4850/foozy/pkg/router/rest"

type ProjectAuthor struct {
	rest.ImplementDTOMessage
	TypProjectAuthor rest.TypeId `dto:"-typeid"`
	UID              int         `dto:"UID" db:"user_id"`
	Username         string      `dto:"Username" db:"username"`
	Avatar           string      `dto:"Avatar" db:"avatar"`
}

type ProjectMessage struct {
	rest.ImplementDTOMessage
	TypProjectMessage rest.TypeId   `dto:"-typeid"`
	Author            ProjectAuthor `dto:"Author"`
	Id                int           `dto:"Id" db:"id"`
	UserId            int           `dto:"UserId" db:"user_id"`
	Name              string        `dto:"Name" db:"name"`
	Description       string        `dto:"Description" db:"description"`
	Error             string        `dto:"Error"`
}

type ProjectLogGroupMessage struct {
	rest.ImplementDTOMessage
	TypProjectLogGroupMessage rest.TypeId `dto:"-typeid"`
	Id                        int         `dto:"Id" db:"id"`
	ProjectId                 int         `dto:"ProjectId" db:"project_id"`
	Name                      string      `dto:"Name" db:"name"`
	Description               string      `dto:"Description" db:"description"`
	Error                     string      `dto:"Error"`
}
