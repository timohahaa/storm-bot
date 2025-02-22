package analytics

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/internal/modules/analytics"
)

func New(conn *pgxpool.Pool) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: analytics.New(conn),
		}
	)

	mux.Get("/week-stat", h.weekStat)

	return mux
}
