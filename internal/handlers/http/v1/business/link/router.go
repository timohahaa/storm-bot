package link

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/internal/modules/link"
)

func New(conn *pgxpool.Pool) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: link.New(conn),
		}
	)

	mux.Post("/", h.create)
	mux.Route("/{link_id}", func(mux chi.Router) {
		mux.Get("/", h.get)
		mux.Delete("/", h.delete)
	})

	return mux
}
