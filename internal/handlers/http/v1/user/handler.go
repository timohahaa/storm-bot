package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/middleware"
	"github.com/timohahaa/gw/internal/modules/user"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

type handler struct {
	mod *user.Module
}

// @Summary	Get user
// @Tags		User
// @Param		user_id	path		string				true	"User ID"
// @Success	200		{object}	user.Model			"User Model"
// @Failure	default	{object}	errors.HTTPError	"Error"
// @Router		/api/v1/users/{user_id}/ [get]
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	var (
		userId, err = uuid.Parse(chi.URLParam(r, "user_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	user, err := h.mod.Get(r.Context(), userId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, user)
}

// @Summary	Update user
// @Tags		User
// @Param		user_id			path		string				true	"User ID"
// @Param		UpdateParams	body		user.UpdateParams	true	"Update Params"
// @Success	200				{object}	user.Model			"User Model"
// @Failure	default			{object}	errors.HTTPError	"Error"
// @Router		/api/v1/users/{user_id}/ [put]
func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	var (
		userId, err = uuid.Parse(chi.URLParam(r, "user_id"))
		form        user.UpdateParams
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	sess := r.Context().Value(middleware.CurrentUser).(*middleware.UserSession)
	if sess == nil || sess.UserID != userId {
		render.Error(w, r, errors.Get(errors.AccessDenied))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		render.Error(w, r, err)
		return
	}

	if invParams := validate.Struct(&form); len(invParams) != 0 {
		render.Error(w, r, invParams)
		return
	}

	user, err := h.mod.Update(r.Context(), userId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, user)
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		userId, err = uuid.Parse(chi.URLParam(r, "user_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.Delete(r.Context(), userId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
