package location

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/location"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

type handler struct {
	mod *location.Module
}

//	@Summary	Create location
//	@Tags		Business Location
//	@Param		business_id		path		string					true	"Business ID"
//	@Param		CreateParams	body		location.CreateParams	true	"Create Params"
//	@Success	200				{object}	location.Model			"Response"
//	@Failure	default			{object}	errors.HTTPError		"Error"
//	@Router		/api/v1/businesses/{business_id}/locations/ [post]
func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	var (
		form            location.CreateParams
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
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

	location, err := h.mod.Create(r.Context(), businessId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, location)
}

//	@Summary	Get location
//	@Tags		Business Location
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		location_id	path		string				true	"Location ID"
//	@Success	200			{object}	location.Model		"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/locations/{location_id}/ [get]
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	var (
		locationId, err = uuid.Parse(chi.URLParam(r, "location_id"))
		businessId      uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	if err != nil {
		render.Error(w, r, err)
		return
	}

	location, err := h.mod.Get(r.Context(), locationId, businessId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, location)
}

//	@Summary	Update location
//	@Tags		Business Location
//	@Param		business_id		path		string					true	"Business ID"
//	@Param		location_id		path		string					true	"Location ID"
//	@Param		UpdateParams	body		location.UpdateParams	true	"Update Params"
//	@Success	200				{object}	location.Model			"Response"
//	@Failure	default			{object}	errors.HTTPError		"Error"
//	@Router		/api/v1/businesses/{business_id}/locations/{location_id}/ [put]
func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	var (
		locationId, err = uuid.Parse(chi.URLParam(r, "location_id"))
		businessId      uuid.UUID
		form            location.UpdateParams
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
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

	location, err := h.mod.Update(r.Context(), locationId, businessId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, location)
}

//	@Summary	Delete location
//	@Tags		Business Location
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		location_id	path		string				true	"Location ID"
//	@Success	204			{object}	nil					"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/locations/{location_id}/ [delete]
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		locationId, err = uuid.Parse(chi.URLParam(r, "location_id"))
		businessId      uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.Delete(r.Context(), locationId, businessId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
