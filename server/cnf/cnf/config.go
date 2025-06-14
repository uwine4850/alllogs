package cnf

import (
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

var DTO = rest.NewDTO()

var DATABASE_ARGS = database.DbArgs{
	Username:     "root",
	Password:     "1111",
	Host:         "localhost", // mysql for docker
	Port:         "3306",
	DatabaseName: "alllogs",
}

const (
	DBT_AUTH              = "auth"
	DBT_PROFILE           = "profile"
	DBT_PROJECT           = "project"
	DBT_PROJECT_LOG_GROUP = "project_log_group"
)

const (
	DEFAULT_AVATAR_PATH = "/storage/avatars/default.jpg"
	STORAGE_AVATAR_PATH = "/storage/avatars/"
)

var DatabaseReader interfaces.IReadDatabase
