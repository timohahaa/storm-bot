package pcard

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/card"
	"github.com/timohahaa/gw/internal/modules/pcard"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

var (
	defaultLimit  = 20
	defaultOffset = 0
)

type mod struct {
	pcard *pcard.Module
	card  *card.Module
}

type handler struct {
	mod mod
}

//	@Summary	List business punch cards
//	@Tags		Business Punch Cards
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		limit		query		string				true	"Pagination Limit"
//	@Param		offset		query		string				true	"Pagination Offset"
//	@Success	200			{object}	listReturnValue		"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/punch-cards/ [get]
func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	var (
		limit, offset = defaultLimit, defaultOffset
	)

	if param := r.URL.Query().Get("limit"); param != "" {
		v, err := strconv.Atoi(param)
		if err != nil {
			render.Error(w, r, err)
			return
		}
		limit = v
	}

	if param := r.URL.Query().Get("offset"); param != "" {
		v, err := strconv.Atoi(param)
		if err != nil {
			render.Error(w, r, err)
			return
		}
		offset = v
	}

	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	cards, err := h.mod.pcard.List(r.Context(), businessId, limit, offset)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	totalCards, err := h.mod.pcard.Total(r.Context(), businessId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSONWithMeta(w, cards, render.Pagination{
		Limit:  limit,
		Offset: offset,
		Total:  totalCards,
	})
}

//	@Summary	Create business punch card
//	@Tags		Business Punch Cards
//	@Param		business_id		path		string				true	"Business ID"
//	@Param		CreateParams	body		pcard.CreateParams	true	"Create Params"
//	@Success	200				{object}	string				"Response"
//	@Failure	default			{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/punch-cards/ [post]
func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	var (
		form            pcard.CreateParams
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

	pcard, err := h.mod.pcard.Create(r.Context(), businessId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, pcard)
}

//	@Summary	Get business punch card
//	@Tags		Business Punch Cards
//	@Param		business_id		path		string				true	"Business ID"
//	@Param		punch_card_id	path		string				true	"Punch Card ID"
//	@Success	200				{object}	pcard.Model			"Response"
//	@Failure	default			{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/punch-cards/{punch_card_id}/ [get]
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	var (
		pcardId, err = uuid.Parse(chi.URLParam(r, "punch_card_id"))
		businessId   uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if businessId, err = uuid.Parse(chi.URLParam(r, "business_id")); err != nil {
		render.Error(w, r, err)
		return
	}

	pcard, err := h.mod.pcard.Get(r.Context(), pcardId, businessId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, pcard)
}

//	@Summary	Update business punch card
//	@Tags		Business Punch Cards
//	@Param		business_id		path		string				true	"Business ID"
//	@Param		punch_card_id	path		string				true	"Punch Card ID"
//	@Param		UpdateParams	body		pcard.UpdateParams	true	"Update Params"
//	@Success	200				{object}	pcard.Model			"Response"
//	@Failure	default			{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/punch-cards/{punch_card_id}/ [put]
func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	var (
		pcardId, err = uuid.Parse(chi.URLParam(r, "punch_card_id"))
		businessId   uuid.UUID
		form         pcard.UpdateParams
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if businessId, err = uuid.Parse(chi.URLParam(r, "business_id")); err != nil {
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

	pcard, err := h.mod.pcard.Update(r.Context(), pcardId, businessId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, pcard)
}

//	@Summary	Delete business punch card
//	@Tags		Business Punch Cards
//	@Param		business_id		path		string				true	"Business ID"
//	@Param		punch_card_id	path		string				true	"Punch Card ID"
//	@Success	204				{object}	nil					"Response"
//	@Failure	default			{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/punch-cards/{punch_card_id}/ [delete]
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		pcardId, err = uuid.Parse(chi.URLParam(r, "punch_card_id"))
		businessId   uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if businessId, err = uuid.Parse(chi.URLParam(r, "business_id")); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.pcard.Delete(r.Context(), pcardId, businessId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//	@Summary	Stamp a punch card
//	@Tags		Business Punch Cards
//	@Param		business_id		path		string				true	"Business ID"
//	@Param		punch_card_id	path		string				true	"Punch Card ID"
//	@Param		customer_id		query		string				true	"User ID"
//	@Success	204				{object}	nil					"Response"
//	@Failure	default			{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/punch-cards/{punch_card_id}/stamp [post]
func (h *handler) stamp(w http.ResponseWriter, r *http.Request) {
	var (
		pcardId, err       = uuid.Parse(chi.URLParam(r, "punch_card_id"))
		businessId, userId uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if businessId, err = uuid.Parse(chi.URLParam(r, "business_id")); err != nil {
		render.Error(w, r, err)
		return
	}

	if userId, err = uuid.Parse(r.URL.Query().Get("user_id")); err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.card.StampByPunchCardId(r.Context(), userId, pcardId, businessId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// added in order for swagger to determine response JSON schema
type listReturnValue struct {
	Data []pcard.Model     `json:"data"`
	Meta render.Pagination `json:"meta"`
}
