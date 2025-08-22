package rprofile

import (
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/uwine4850/alllogs/api"
	"github.com/uwine4850/alllogs/api/apiform"
	"github.com/uwine4850/alllogs/api/permissions/profileperm"
	"github.com/uwine4850/alllogs/cnf/cnf"
	qb "github.com/uwine4850/foozy/pkg/database/querybuld"
	"github.com/uwine4850/foozy/pkg/interfaces"
	"github.com/uwine4850/foozy/pkg/router/form"
	"github.com/uwine4850/foozy/pkg/router/rest"
)

type MsgProfileUpdate struct {
	rest.ImplementDTOMessage
	TypProfileUpdateMessage rest.TypeId   `dto:"-typeid"`
	UID                     int           `dto:"UID"`
	Description             string        `dto:"Description"`
	Avatar                  form.FormFile `dto:"Avatar"`
	OldAvatarPath           string        `dto:"OldAvatarPath"`
	DelAvatar               bool          `dto:"DelAvatar"`
}

type UpdateForm struct {
	UID           int           `form:"UID" empty:"-err"`
	Description   string        `form:"Description" nil:"-skip" empty:"-err"`
	Avatar        form.FormFile `form:"Avatar" nil:"-skip" empty:"-err"`
	OldAvatarPath string        `form:"OldAvatarPath" nil:"-skip" empty:"-err"`
	DelAvatar     string        `form:"DelAvatar" empty:"-err"`
}

func Update(w http.ResponseWriter, r *http.Request, manager interfaces.Manager) error {
	updateForm := UpdateForm{}
	if err := apiform.ParseAndFill(r, &updateForm); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())

	}
	if err := profileperm.ProfilePermission(manager, updateForm.UID, "no access for user profile updates"); err != nil {
		return err
	}

	var oldRelativeAvatarPath string
	if updateForm.OldAvatarPath != "" {
		if updateForm.OldAvatarPath == cnf.DEFAULT_AVATAR_PATH {
			oldRelativeAvatarPath = updateForm.OldAvatarPath
		} else {
			oldRelativeAvatarPath = filepath.Join("../client/public/", updateForm.OldAvatarPath)
		}
	}

	var newAvatarPath string
	if err := updateAvatar(&newAvatarPath, oldRelativeAvatarPath, &updateForm, manager); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}

	var description string
	if updateForm.Description != "" {
		description = updateForm.Description
	}
	if err := saveUpdate(description, newAvatarPath, &updateForm); err != nil {
		return api.NewServerError(http.StatusInternalServerError, err.Error())
	}
	api.SendBeseResponse(w, true, nil)
	return nil
}

func updateAvatar(newAvatarPath *string, oldRelativeAvatarPath string, updateForm *UpdateForm, manager interfaces.Manager) error {
	isDelAvatar, err := strconv.ParseBool(updateForm.DelAvatar)
	if err != nil {
		return err
	}
	if !isDelAvatar && !reflect.DeepEqual(updateForm.Avatar, form.FormFile{}) && oldRelativeAvatarPath != "" {
		avatar := updateForm.Avatar
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
	newQB.Update(cnf.DBT_PROFILE, updateArgs).Where(qb.Compare("user_id", qb.EQUAL, updateForm.UID))
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
