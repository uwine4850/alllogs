package apiform

import (
	"net/http"

	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
)

func ParseAndFill[T any](r *http.Request, out *T) error {
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		return err
	}
	if err := mapper.FillStructFromForm(frm, out); err != nil {
		return err
	}
	return nil
}
