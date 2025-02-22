package card

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/card"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

var (
	defaultLimit  = 20
	defaultOffset = 0
)

type handler struct {
	mod *card.Module
}

//	@Summary	Create card
//	@Tags		Customer Cards
//	@Param		customer_id		path		string				true	"Customer ID"
//	@Param		CreateParams	body		card.CreateParams	true	"Create Params"
//	@Success	200				{object}	card.Model			"Response"
//	@Failure	default			{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/customers/{customer_id}/cards/ [post]
func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	var (
		form            card.CreateParams
		customerId, err = uuid.Parse(chi.URLParam(r, "customer_id"))
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

	card, err := h.mod.Create(r.Context(), customerId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, card)
}

//	@Summary	Get card
//	@Tags		Customer Cards
//	@Param		card_id		path		string				true	"Card ID"
//	@Param		customer_id	path		string				true	"Customer ID"
//	@Success	200			{object}	card.Model			"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/customers/{customer_id}/cards/{card_id}/ [get]
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	var (
		cardId, err = uuid.Parse(chi.URLParam(r, "card_id"))
		customerId  uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	customerId, err = uuid.Parse(chi.URLParam(r, "customer_id"))
	if err != nil {
		render.Error(w, r, err)
		return
	}

	card, err := h.mod.Get(r.Context(), cardId, customerId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, card)
}

//	@Summary	List cards
//	@Tags		Customer Cards
//	@Param		limit		query		string				true	"Customer ID"
//	@Param		offset		query		string				true	"Customer ID"
//	@Param		customer_id	path		string				true	"Customer ID"
//	@Success	200			{object}	[]card.Model		"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/customers/{customer_id}/cards/ [get]
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
		customerId, err = uuid.Parse(chi.URLParam(r, "customer_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	cards, err := h.mod.List(r.Context(), customerId, limit, offset)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	totalCards, err := h.mod.Total(r.Context(), customerId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSONWithMeta(w, cards, render.Pagination{
		Limit:  limit,
		Offset: offset,
		Total:  totalCards,
	})

	render.JSON(w, cards)
}

//	@Summary	Delete card
//	@Tags		Customer Cards
//	@Param		card_id		path		string				true	"Card ID"
//	@Param		customer_id	path		string				true	"Customer ID"
//	@Success	204			{object}	nil					"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/customers/{customer_id}/cards/{card_id}/ [delete]
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		cardId, err = uuid.Parse(chi.URLParam(r, "card_id"))
		customerId  uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	customerId, err = uuid.Parse(chi.URLParam(r, "customer_id"))
	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.Delete(r.Context(), cardId, customerId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
