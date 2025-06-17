package rproject

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type LogItemsView struct {
	object.AllView
	Database       interfaces.IReadDatabase
	LogGroupSlugId string
	StartSlugId    string
	CountSlug      string
}

func (v *LogItemsView) OnError(w http.ResponseWriter, r *http.Request, m interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *LogItemsView) Object(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	logGroupId, ok := manager.OneTimeData().GetSlugParams(v.LogGroupSlugId)
	if !ok {
		return nil, object.ErrNoSlug{SlugName: v.StartSlugId}
	}

	var startValue any
	startSlug, ok := manager.OneTimeData().GetSlugParams(v.StartSlugId)
	if !ok {
		return nil, object.ErrNoSlug{SlugName: v.StartSlugId}
	}
	if startSlug == "-1" {
		startValue = qb.SQ(true,
			qb.NewNoDbQB().SelectFrom("MAX(id)", cnf.DBT_LOG_ITEM).
				Where(qb.Compare("log_group_id", qb.EQUAL, logGroupId)),
		)
	} else {
		startValue = startSlug
	}

	countSlug, ok := manager.OneTimeData().GetSlugParams(v.CountSlug)
	if !ok {
		return nil, object.ErrNoSlug{SlugName: v.CountSlug}
	}
	count, err := strconv.Atoi(countSlug)
	if err != nil {
		return nil, err
	}
	newQB := qb.NewSyncQB(v.Database.SyncQ())
	newQB.SelectFrom("*", v.TableName).
		Where(
			qb.Compare("log_group_id", qb.EQUAL, logGroupId), qb.AND,
			qb.Compare("id", "<=", startValue)).OrderBy(qb.DESC("id")).Limit(count).Merge()
	res, err := newQB.Query()
	if err != nil {
		return nil, err
	}
	filledObjects, err := v.fillObjects(res)
	if err != nil {
		return nil, err
	}

	return object.Context{v.Name: filledObjects}, nil
}

func (v *LogItemsView) fillObjects(objects []map[string]interface{}) ([]interface{}, error) {
	if v.FillStruct == nil {
		panic("the FillStruct field must not be nil")
	}
	if typeopr.IsPointer(v.FillStruct) {
		return nil, typeopr.ErrValueIsPointer{Value: "FillStruct"}
	}
	var objectsStruct []interface{}
	for i := 0; i < len(objects); i++ {
		fillType := reflect.TypeOf(v.FillStruct)
		value := reflect.New(fillType).Elem()
		err := mapper.FillStructFromDb(&value, &objects[i])
		if err != nil {
			return nil, err
		}
		objectsStruct = append(objectsStruct, value.Interface())
	}
	return objectsStruct, nil
}

func LogItemsObjectView(database interfaces.IReadDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	d := object.NewViewMysqlDatabase(database)
	view := object.JsonAllTemplateView{
		View: &LogItemsView{
			AllView: object.AllView{
				Name:       "logs",
				TableName:  cnf.DBT_LOG_ITEM,
				Database:   d,
				FillStruct: LogItemPayload{},
			},
			Database:       database,
			LogGroupSlugId: "logGroupId",
			StartSlugId:    "start",
			CountSlug:      "count",
		},
		DTO:     cnf.DTO,
		Message: LogItemPayload{},
	}
	return view.Call
}
