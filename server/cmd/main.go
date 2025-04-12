package main

import (
	"errors"
	"net/http"

	"github.com/rs/cors"
	"github.com/uwine4850/alllogs/cnf/cnf"
	initcnf "github.com/uwine4850/alllogs/cnf/init_cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/alllogs/rest/rauth"
	"github.com/uwine4850/alllogs/routes"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/builtin/bglobalflow"
	"github.com/uwine4850/foozy/pkg/builtin/builtin_mddl"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/manager"
	"github.com/uwine4850/foozy/pkg/router/middlewares"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
	"github.com/uwine4850/foozy/pkg/server"
	"github.com/uwine4850/foozy/pkg/server/globalflow"
)

func main() {
	initcnf.InitCnf()
	gDB := database.NewDatabase(cnf.DATABASE_ARGS)
	if err := gDB.Connect(); err != nil {
		panic(err)
	}
	defer gDB.Close()

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
	newManager.Key().Generate32BytesKeys()
	newMiddleware := middlewares.NewMiddleware()
	newMiddleware.HandlerMddl(0, builtin_mddl.AuthJWT(
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (string, error) {
			tokenStr := r.Header.Get("Authorization")
			if tokenStr == "" {
				return "", nil
			}
			return tokenStr, nil
		},
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, token string) error {
			middlewares.SkipNextPage(manager.OneTimeData())
			rauth.SendLoginResponse(w, token, "")()
			return nil
		},
		func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
			middlewares.SkipNextPage(manager.OneTimeData())
			rauth.SendLoginResponse(w, "", err.Error())()
		},
	))
	newRouter := router.NewRouter(newManager)
	newRouter.SetMiddleware(newMiddleware)

	newRouter.AddHandlerSet(routes.Get())

	serv := server.NewServer(":8000", newRouter, &cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	go func() {
		gf := globalflow.NewGlobalFlow(1000)
		gf.AddNotWaitTask(bglobalflow.KeyUpdater(10))
		gf.Run(newManager)
	}()

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
