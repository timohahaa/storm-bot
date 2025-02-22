package link

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/link"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

type handler struct {
	mod *link.Module
}

//	@Summary	Create link
//	@Tags		Business Links
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		type		query		string				true	"Link Type (one-time, permanent)"
//	@Success	200			{object}	link.Model			"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/links/ [post]
func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
		linkType        = r.URL.Query().Get("type")
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if invParams := validate.Var(linkType, "oneof=one-time permanent", "LinkType"); len(invParams) != 0 {
		render.Error(w, r, invParams)
		return
	}

	link, err := h.mod.Create(r.Context(), businessId, link.LinkType(linkType))
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, link)
}

//	@Summary	Create link
//	@Tags		Business Links
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		link_id		path		string				true	"Link ID"
//	@Success	200			{object}	link.Model			"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/links/{link_id}/ [get]
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	var (
		linkId, err = uuid.Parse(chi.URLParam(r, "link_id"))
		businessId  uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if businessId, err = uuid.Parse(chi.URLParam(r, "business_id")); err != nil {
		render.Error(w, r, err)
		return
	}

	link, err := h.mod.Get(r.Context(), linkId, businessId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, link)
}

//	@Summary	Delete link
//	@Tags		Business Links
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		link_id		path		string				true	"Link ID"
//	@Success	204			{object}	nil					"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/links/{link_id}/ [delete]
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		linkId, err = uuid.Parse(chi.URLParam(r, "link_id"))
		businessId  uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if businessId, err = uuid.Parse(chi.URLParam(r, "business_id")); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.Delete(r.Context(), linkId, businessId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
