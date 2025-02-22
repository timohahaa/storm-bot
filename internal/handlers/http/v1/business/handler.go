package business

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/middleware"
	"github.com/timohahaa/gw/internal/modules/business"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

type handler struct {
	mod *business.Module
}

//	@Summary	Create Business
//	@Tags		Business
//	@Param		CreateParams	body		business.CreateParams	true	"Create Params"
//	@Success	200				{object}	business.Model			"Business Model"
//	@Failure	default			{object}	errors.HTTPError		"Error"
//	@Router		/api/v1/businesses/ [post]
func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	var (
		form business.CreateParams
	)

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		render.Error(w, r, err)
		return
	}

	form.UserID = r.Context().Value(middleware.CurrentUser).(*middleware.UserSession).UserID

	if invParams := validate.Struct(&form); len(invParams) != 0 {
		render.Error(w, r, invParams)
		return
	}

	business, err := h.mod.Create(r.Context(), form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, business)
}

//	@Summary	Get business
//	@Tags		Business
//	@Param		business_id	path		string				true	"Business ID"
//	@Success	200			{object}	business.Model		"Business Model"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/ [get]
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	business, err := h.mod.Get(r.Context(), businessId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, business)
}

//	@Summary	Update business
//	@Tags		Business
//	@Param		business_id		path		string					true	"Business ID"
//	@Param		UpdateParams	body		business.UpdateParams	true	"Update Params"
//	@Success	200				{object}	business.Model			"Business Model"
//	@Failure	default			{object}	errors.HTTPError		"Error"
//	@Router		/api/v1/businesses/{business_id}/ [put]
func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
		form            business.UpdateParams
	)

	if err != nil {
		render.Error(w, r, err)
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

	business, err := h.mod.Update(r.Context(), businessId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, business)
}

//	@Summary	Delete business
//	@Tags		Business
//	@Param		business_id	path		string				true	"Business ID"
//	@Success	204			{object}	nil					"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/ [delete]
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.Delete(r.Context(), businessId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
