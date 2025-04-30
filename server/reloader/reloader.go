package main

import (
	initcnf "github.com/uwine4850/alllogs/cnf/init_cnf"
	"github.com/uwine4850/foozy/pkg/server/livereload"
)

func main() {
	initcnf.InitCnf()
	wrt := livereload.NewWiretap()
	wrt.SetDirs([]string{"cmd", "cnf", "reloader", "mydto", "api", "middlewares"})
	wrt.SetExcludeDirs([]string{"cnf/log"})
	reload := livereload.NewReloader("cmd/main.go", wrt)
	if err := reload.Start(); err != nil {
		panic(err)
	}
}
