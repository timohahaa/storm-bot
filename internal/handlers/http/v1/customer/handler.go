package customer

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/customer"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

type handler struct {
	mod *customer.Module
}

// @Summary	Get customer
// @Tags		Customer
// @Param		customer_id	path		string				true	"Customer ID"
// @Success	200			{object}	customer.Model		"Response"
// @Failure	default		{object}	errors.HTTPError	"Error"
// @Router		/api/v1/customers/{customer_id}/ [get]
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	var (
		customerId, err = uuid.Parse(chi.URLParam(r, "customer_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	customer, err := h.mod.Get(r.Context(), customerId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, customer)
}

// @Summary	Update customer
// @Tags		Customer
// @Param		customer_id		path		string					true	"Customer ID"
// @Param		UpdateParams	body		customer.UpdateParams	true	"Update Params"
// @Success	200				{object}	customer.Model			"Response"
// @Failure	default			{object}	errors.HTTPError		"Error"
// @Router		/api/v1/customers/{customer_id}/ [put]
func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	var (
		customerId, err = uuid.Parse(chi.URLParam(r, "customer_id"))
		form            customer.UpdateParams
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

	customer, err := h.mod.Update(r.Context(), customerId, form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, customer)
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		customerId, err = uuid.Parse(chi.URLParam(r, "customer_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.Delete(r.Context(), customerId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
