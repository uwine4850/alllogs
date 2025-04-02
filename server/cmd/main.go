package main

import (
	"errors"
	"net/http"

	"github.com/rs/cors"
	"github.com/uwine4850/alllogs/cnf/cnf"
	initcnf "github.com/uwine4850/alllogs/cnf/init_cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/alllogs/routes"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/manager"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
	"github.com/uwine4850/foozy/pkg/server"
)

func main() {
	initcnf.InitCnf()
	if err := createAuthTable(); err != nil {
		panic(err)
	}

	mydto.SetUpMessages(mydto.DTO)
	if err := mydto.DTO.Generate(); err != nil {
		panic(err)
	}

	render, err := tmlengine.NewRender()
	if err != nil {
		panic(err)
	}
	newManager := manager.NewManager(render)
	newRouter := router.NewRouter(newManager)

	newRouter.AddHandlerSet(routes.Get())

	serv := server.NewServer(":8000", newRouter, &cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	err = serv.Start()
	if err != nil && !errors.Is(http.ErrServerClosed, err) {
		panic(err)
	}
}

func createAuthTable() error {
	db := database.NewDatabase(cnf.DATABASE_ARGS)
	if err := db.Connect(); err != nil {
		return err
	}

	if err := auth.CreateAuthTable(db); err != nil {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
