package cnf

import "github.com/uwine4850/foozy/pkg/database"

var DATABASE_ARGS = database.DbArgs{
	Username:     "root",
	Password:     "1111",
	Host:         "localhost", // mysql for docker
	Port:         "3306",
	DatabaseName: "alllogs",
}

const (
	DBT_AUTH    = "auth"
	DBT_PROFILE = "profile"
)

const (
	DEFAULT_AVATAR_PATH = "/storage/avatars/default.jpg"
	STORAGE_AVATAR_PATH = "/storage/avatars/"
)
