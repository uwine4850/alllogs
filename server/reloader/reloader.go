package main

import (
	initcnf "github.com/uwine4850/alllogs/cnf/init_cnf"
	"github.com/uwine4850/foozy/pkg/server/livereload"
)

func main() {
	initcnf.InitCnf()
	reload := livereload.NewReload("cmd/main.go", livereload.NewWiretap([]string{"cmd", "cnf", "reloader", "mydto", "rest"}, []string{"cnf/log"}))
	reload.Start()
}
