package analytics

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/analytics"
	"github.com/timohahaa/gw/internal/utils/render"
)

type handler struct {
	mod *analytics.Module
}

//	@Summary	Get week business stats
//	@Tags		Business Analytics
//	@Param		business_id	path		string				true	"Business ID"
//	@Success	200			{object}	analytics.WeekStat	"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/analytics/week-stat [get]
func (h *handler) weekStat(w http.ResponseWriter, r *http.Request) {
	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	stat, err := h.mod.WeekStat(r.Context(), businessId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, stat)
}
