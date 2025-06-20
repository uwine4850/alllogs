package main

import (
	"errors"
	"net/http"

	"github.com/rs/cors"
	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	initcnf "github.com/uwine4850/alllogs/cnf/init_cnf"
	"github.com/uwine4850/alllogs/middlewares/mddlauth"
	"github.com/uwine4850/alllogs/middlewares/securemddl"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/alllogs/routes"
	"github.com/uwine4850/foozy/pkg/builtin/auth"
	"github.com/uwine4850/foozy/pkg/builtin/bglobalflow"
	"github.com/uwine4850/foozy/pkg/config"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/manager"
	"github.com/uwine4850/foozy/pkg/router/middlewares"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
	"github.com/uwine4850/foozy/pkg/server"
	"github.com/uwine4850/foozy/pkg/server/globalflow"
)

func main() {
	initcnf.InitCnf()

	mydto.SetUpMessages(cnf.DTO)
	if err := cnf.DTO.Generate(); err != nil {
		panic(err)
	}

	syncQ := database.NewSyncQueries()
	gDB := database.NewMysqlDatabase(cnf.DATABASE_ARGS, syncQ, database.NewAsyncQueries(syncQ))
	if err := gDB.Open(); err != nil {
		panic(err)
	}
	defer gDB.Close()
	render, err := tmlengine.NewRender()
	if err != nil {
		panic(err)
	}
	newManager := manager.NewManager(manager.NewOneTimeData(), render, manager.NewDatabasePool())
	database.InitDatabasePool(newManager, gDB)

	dbRead, err := newManager.Database().ConnectionPool(config.LoadedConfig().Default.Database.MainConnectionPoolName)
	if err != nil {
		panic(err)
	}
	cnf.DatabaseReader = dbRead
	if err := auth.CreateMysqlAuthTable(dbRead, cnf.DATABASE_ARGS.DatabaseName); err != nil {
		panic(err)
	}

	newManager.Key().Generate32BytesKeys()
	newMiddleware := middlewares.NewMiddlewares()
	newMiddleware.PreMiddleware(0, securemddl.ValidateCSRFToken)
	newMiddleware.PreMiddleware(1, mddlauth.CheckJWT)
	newAdapter := router.NewAdapter(newManager, newMiddleware)
	newAdapter.SetOnErrorFunc(func(w http.ResponseWriter, r *http.Request, err error) {
		var clientErrortarget *api.ClientError
		if errors.As(err, &clientErrortarget) {
			clientError := err.(*api.ClientError)
			api.SendClientError(w, clientError.Code, clientError.Text)
			return
		}
		var serverErrorTarget *api.ServerError
		if errors.As(err, &serverErrorTarget) {
			serverError := err.(*api.ServerError)
			api.SendServerError(w, serverError.Code, serverError.Text)
			return
		}
	})
	newRouter := router.NewRouter(newAdapter)

	newRouter.HandlerSet(routes.Get(gDB))

	serv := server.NewServer(":8000", newRouter, &cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-CSRF-TOKEN"},
		AllowCredentials: true,
	})

	go func() {
		gf := globalflow.NewGlobalFlow(1000)            // 1 sec
		gf.AddNotWaitTask(bglobalflow.KeyUpdater(3600)) // 1 hour
		gf.Run(newManager)
	}()

	err = serv.Start()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
