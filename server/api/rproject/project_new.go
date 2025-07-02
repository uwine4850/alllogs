package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

// TODO: add permission

type MsgProjectAuthor struct {
	rest.ImplementDTOMessage
	TypProjectAuthor rest.TypeId `dto:"-typeid"`
	UID              int         `dto:"UID" db:"user_id"`
	Username         string      `dto:"Username" db:"username"`
	Avatar           string      `dto:"Avatar" db:"avatar"`
}

type MsgProject struct {
	rest.ImplementDTOMessage
	TypProjectMessage rest.TypeId      `dto:"-typeid"`
	Author            MsgProjectAuthor `dto:"Author"`
	Id                int              `dto:"Id" db:"id"`
	UserId            int              `dto:"UserId" db:"user_id"`
	Name              string           `dto:"Name" db:"name"`
	Description       string           `dto:"Description" db:"description"`
	Error             string           `dto:"Error"`
}

type ProjectForm struct {
	Name        string `form:"Name" empty:"-err"`
	Description string `form:"Description" empty:"-err"`
}

func NewProject(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	UID, ok := manager.OneTimeData().GetUserContext("UID")
	if !ok {
		return api.NewServerError(http.StatusInternalServerError, "user ID not found")
	}

	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	var projectForm ProjectForm
	if err := mapper.FillStructFromForm(frm, &projectForm); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}

	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ()).Insert(cnf.DBT_PROJECT,
		map[string]any{
			"user_id": UID, "name": projectForm.Name, "description": projectForm.Description,
		})
	newQB.Merge()
	_, err := newQB.Exec()
	if err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}

func IsProjectAuthor(UID int, projectId int, dbRead interfaces.IReadDatabase) (bool, error) {
	newQB := qb.NewSyncQB(dbRead.SyncQ())
	return qb.SelectExists(newQB, cnf.DBT_PROJECT,
		qb.Compare("id", qb.EQUAL, projectId), qb.AND,
		qb.Compare("user_id", qb.EQUAL, UID))
}

func changeProjectPermissions(projectId int, UID any) (bool, error) {
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.SelectFrom("id", cnf.DBT_PROJECT).Where(
		qb.Compare("id", qb.EQUAL, projectId), qb.AND,
		qb.Compare("user_id", qb.EQUAL, UID),
	).Merge()
	res, err := newQB.Query()
	if err != nil {
		return false, err
	}
	if len(res) != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
