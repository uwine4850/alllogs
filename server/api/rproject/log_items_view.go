package rproject

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/router/rest"
	"github.com/uwine4850/foozy/pkg/typeopr"
)

type LogItemsFilterMessage struct {
	rest.ImplementDTOMessage
	TypLogItemsFilter rest.TypeId `dto:"-typeid"`
	Text              string      `dto:"Text"`
	Type              string      `dto:"Type"`
	Tag               string      `dto:"Tag"`
	DateTime          string      `dto:"DateTime"`
}

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
	singleQueryMap := parseSingleQuery(r.URL.Query())
	var logItemsFilterMessage LogItemsFilterMessage
	if err := mapper.JsonToDTOMessage(singleQueryMap, cnf.DTO, &logItemsFilterMessage); err != nil {
		return nil, err
	}
	logGroupId, ok := manager.OneTimeData().GetSlugParams(v.LogGroupSlugId)
	if !ok {
		return nil, object.ErrNoSlug{SlugName: v.StartSlugId}
	}

	startSlug, ok := manager.OneTimeData().GetSlugParams(v.StartSlugId)
	if !ok {
		return nil, object.ErrNoSlug{SlugName: v.StartSlugId}
	}
	filterArgsValue := filterArgs(&logItemsFilterMessage)
	whereArgsValue := whereArgs(startSlug, logGroupId, filterArgsValue)

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
			whereArgsValue...,
		).OrderBy(qb.DESC("id")).Limit(count).Merge()
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

func parseSingleQuery(values url.Values) map[string]interface{} {
	out := map[string]interface{}{}
	for key, sliceValue := range values {
		if len(sliceValue) > 0 {
			out[key] = sliceValue[0]
		}
	}
	return out
}

func filterArgs(logItemsFilterMessage *LogItemsFilterMessage) []any {
	filterValue := []any{}
	if logItemsFilterMessage.Text != "" {
		typeText := qb.NoArgsCompare("text", qb.LIKE, "'%"+logItemsFilterMessage.Text+"%'")
		filterValue = append(filterValue, typeText)
	}
	if logItemsFilterMessage.Type != "" {
		typeFilter := qb.Compare("type", qb.EQUAL, logItemsFilterMessage.Type)
		if len(filterValue) != 0 {
			filterValue = append(filterValue, qb.AND)
		}
		filterValue = append(filterValue, typeFilter)
	}
	if logItemsFilterMessage.Tag != "" {
		tagFilter := qb.NoArgsCompare("tag", qb.LIKE, "'%"+logItemsFilterMessage.Tag+"%'")
		if len(filterValue) != 0 {
			filterValue = append(filterValue, qb.AND)
		}
		filterValue = append(filterValue, tagFilter)
	}
	if logItemsFilterMessage.DateTime != "" {
		dateTimeFilter := qb.Compare("DATE(datetime)", qb.EQUAL, logItemsFilterMessage.DateTime)
		if len(filterValue) != 0 {
			filterValue = append(filterValue, qb.AND)
		}
		filterValue = append(filterValue, dateTimeFilter)
	}
	return filterValue
}

func whereArgs(startSlug string, logGroupId string, filterValue []any) []any {
	whereArgs := []any{}
	qq := qb.Compare("log_group_id", qb.EQUAL, logGroupId)
	whereArgs = append(whereArgs, qq)

	if startSlug == "-1" {
		maxIdSQ := qb.SQ(true,
			qb.NewNoDbQB().SelectFrom("MAX(id)", cnf.DBT_LOG_ITEM).
				Where(qb.Compare("log_group_id", qb.EQUAL, logGroupId)),
		)
		idCompare := qb.Compare("id", "<=", maxIdSQ)
		whereArgs = append(whereArgs, qb.AND)
		whereArgs = append(whereArgs, idCompare)
	} else {
		idCompare := qb.Compare("id", "<=", startSlug)
		whereArgs = append(whereArgs, qb.AND)
		whereArgs = append(whereArgs, idCompare)
	}
	if len(filterValue) != 0 {
		whereArgs = append(whereArgs, qb.AND)
		whereArgs = append(whereArgs, filterValue...)
	}
	return whereArgs
}
