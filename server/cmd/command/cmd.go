package main

import (
	initcnf "github.com/uwine4850/alllogs/cnf/init_cnf"
	"github.com/uwine4850/foozy/pkg/cmd"
)

func main() {
	initcnf.InitCnf()
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
