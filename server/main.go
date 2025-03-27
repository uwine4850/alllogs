package main

import (
	"errors"
	"net/http"

	initcnf "github.com/uwine4850/alllogs/cnf/init_cnf"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/manager"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
	"github.com/uwine4850/foozy/pkg/server"
)

func main() {
	initcnf.InitCnf()
	render, err := tmlengine.NewRender()
	if err != nil {
		panic(err)
	}
	newManager := manager.NewManager(render)
	newRouter := router.NewRouter(newManager)
	newRouter.Get("/home", func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
		return func() { w.Write([]byte("HELLO WORLD!")) }
	})
	serv := server.NewServer(":8000", newRouter, nil)
	err = serv.Start()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		panic(err)
	}
}
