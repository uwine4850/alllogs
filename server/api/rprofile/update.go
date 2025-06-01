package rprofile

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/mapper"
	"github.com/uwine4850/foozy/pkg/router/form"
)

type UpdateForm struct {
	PID           []string        `form:"PID"`
	Description   []string        `form:"Description" nil:"-skip"`
	Avatar        []form.FormFile `form:"Avatar" nil:"-skip"`
	OldAvatarPath []string        `form:"OldAvatarPath" nil:"-skip"`
	DelAvatar     []string        `form:"DelAvatar"`
}

func Update(w http.ResponseWriter, r *http.Request, manager interfaces.IManager) error {
	frm := form.NewForm(r)
	if err := frm.Parse(); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	updateForm := UpdateForm{}
	if err := mapper.FillStructFromForm(frm, &updateForm); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	var oldRelativeAvatarPath string
	if updateForm.OldAvatarPath != nil {
		if updateForm.OldAvatarPath[0] == cnf.DEFAULT_AVATAR_PATH {
			oldRelativeAvatarPath = updateForm.OldAvatarPath[0]
		} else {
			oldRelativeAvatarPath = filepath.Join("../client/public/", updateForm.OldAvatarPath[0])
		}
	}

	var newAvatarPath string
	if err := updateAvatar(&newAvatarPath, oldRelativeAvatarPath, &updateForm, manager); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}

	var description string
	if updateForm.Description != nil {
		description = updateForm.Description[0]
	}
	if err := saveUpdate(description, newAvatarPath, &updateForm); err != nil {
		api.SendBeseResponse(w, false, err)
		return nil
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}

func updateAvatar(newAvatarPath *string, oldRelativeAvatarPath string, updateForm *UpdateForm, manager interfaces.IManager) error {
	isDelAvatar, err := strconv.ParseBool(updateForm.DelAvatar[0])
	if err != nil {
		return err
	}
	if !isDelAvatar && updateForm.Avatar != nil && oldRelativeAvatarPath != "nil" {
		avatar := updateForm.Avatar[0]
		if oldRelativeAvatarPath == cnf.DEFAULT_AVATAR_PATH {
			if err := form.SaveFile(avatar.Header, "../client/public/storage/avatars", newAvatarPath, manager); err != nil {
				return err
			}
		} else {
			err := form.ReplaceFile(oldRelativeAvatarPath, avatar.Header, "../client/public/storage/avatars", newAvatarPath, manager)
			if err != nil {
				return err
			}
		}
	}
	if isDelAvatar && oldRelativeAvatarPath != "" {
		if oldRelativeAvatarPath != cnf.DEFAULT_AVATAR_PATH {
			err := os.Remove(oldRelativeAvatarPath)
			if err != nil {
				return err
			}
			*newAvatarPath = cnf.DEFAULT_AVATAR_PATH
		}
	}
	return nil
}

func saveUpdate(description string, newAvatarPath string, updateForm *UpdateForm) error {
	updateArgs := map[string]any{"description": description}
	if newAvatarPath != "" {
		updateArgs["avatar"] = filepath.Join(cnf.STORAGE_AVATAR_PATH, filepath.Base(newAvatarPath))
	}
	newQB := qb.NewSyncQB(cnf.DatabaseReader.SyncQ())
	newQB.Update(cnf.DBT_PROFILE, updateArgs).Where(qb.Compare("id", qb.EQUAL, updateForm.PID[0]))
	newQB.Merge()
	_, err := newQB.Exec()
	if err != nil {
		if newAvatarPath != "" {
			err := os.Remove(newAvatarPath)
			if err != nil {
				return err
			}
		}
		return err
	}
	return nil
}
