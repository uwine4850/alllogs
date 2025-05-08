package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/alllogs/mydto"
	"github.com/uwine4850/foozy/pkg/database"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router"
	"github.com/uwine4850/foozy/pkg/router/object"
)

type LogGroupView struct {
	object.MultipleObjectView
}

func (v *LogGroupView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	msg := mydto.ProjectLogGroupMessage{}
	msg.Error = err.Error()
	out := map[string]any{"log": msg}
	router.SendJson(out, w)
}

func (v *LogGroupView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func LogGroupObjectView() func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) func() {
	db := database.NewDatabase(cnf.DATABASE_ARGS)
	if err := db.Connect(); err != nil {
		panic(err)
	}
	view := object.JsonMultipleObjectTemplateView{
		View: &LogGroupView{
			object.MultipleObjectView{
				DB: db,
				MultipleObjects: []object.MultipleObject{
					{
						Name:       "project",
						TaleName:   cnf.DBT_PROJECT,
						SlugName:   "projID",
						SlugField:  "id",
						FillStruct: mydto.ProjectMessage{},
					},
					{
						Name:       "log",
						TaleName:   cnf.DBT_PROJECT_LOG_GROUP,
						SlugName:   "logID",
						SlugField:  "id",
						FillStruct: mydto.ProjectLogGroupMessage{},
					},
				},
			},
		},
		DTO: mydto.DTO,
		Messages: map[string]irest.IMessage{
			"project": mydto.ProjectMessage{},
			"log":     mydto.ProjectLogGroupMessage{},
		},
	}
	// view.OnMessageFilled(func(message any, manager interfaces.IManager) error {
	// 	msg, ok := message.(*mydto.ProjectMessage)
	// 	if !ok {
	// 		return errors.New(("error converting a filled message"))
	// 	}
	// 	_db, ok := manager.OneTimeData().GetUserContext(namelib.OBJECT.OBJECT_DB)
	// 	if !ok {
	// 		return errors.New(("error retrieving database from object"))
	// 	}
	// 	db := _db.(*database.Database)
	// 	newQB := qb.NewSyncQB(db.SyncQ())
	// 	newQB.SelectFrom("auth.username, profile.avatar, profile.id as pid", cnf.DBT_AUTH).
	// 		InnerJoin(cnf.DBT_PROFILE, qb.NoArgsCompare("user_id", qb.EQUAL, "auth.id")).
	// 		Where(qb.Compare("auth.id", qb.EQUAL, msg.UserId))
	// 	newQB.Merge()
	// 	res, err := newQB.Query()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	var author []mydto.ProjectAuthor
	// 	dbMapper := dbmapper.NewMapper(res, typeopr.Ptr{}.New(&author))
	// 	if err := dbMapper.Fill(); err != nil {
	// 		return err
	// 	}
	// 	msg.Author = author[0]
	// 	return nil
	// })
	return view.Call
}
