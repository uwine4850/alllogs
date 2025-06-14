package rproject

import (
	"net/http"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/interfaces/irest"
	"github.com/uwine4850/foozy/pkg/router/object"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type ProjectLogGroupMessage struct {
	rest.ImplementDTOMessage
	TypProjectLogGroupMessage rest.TypeId `dto:"-typeid"`
	Id                        int         `dto:"Id" db:"id"`
	ProjectId                 int         `dto:"ProjectId" db:"project_id"`
	Name                      string      `dto:"Name" db:"name"`
	Description               string      `dto:"Description" db:"description"`
	Error                     string      `dto:"Error"`
}

type LogGroupView struct {
	object.MultipleObjectView
}

func (v *LogGroupView) OnError(w http.ResponseWriter, r *http.Request, manager interfaces.IManager, err error) {
	api.SendServerError(w, http.StatusInternalServerError, err.Error())
}

func (v *LogGroupView) Context(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) (object.Context, error) {
	return object.Context{}, nil
}

func LogGroupObjectView(database object.IViewDatabase) func(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	view := object.JsonMultipleObjectTemplateView{
		View: &LogGroupView{
			object.MultipleObjectView{
				Database: database,
				MultipleObjects: []object.MultipleObject{
					{
						Name:       "project",
						TaleName:   cnf.DBT_PROJECT,
						SlugName:   "projID",
						SlugField:  "id",
						FillStruct: ProjectMessage{},
					},
					{
						Name:       "log",
						TaleName:   cnf.DBT_PROJECT_LOG_GROUP,
						SlugName:   "logID",
						SlugField:  "id",
						FillStruct: ProjectLogGroupMessage{},
					},
				},
			},
		},
		DTO: cnf.DTO,
		Messages: map[string]irest.IMessage{
			"project": ProjectMessage{},
			"log":     ProjectLogGroupMessage{},
		},
	}
	return view.Call
}
