package location

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/internal/modules/location"
)

func New(conn *pgxpool.Pool, mc *memcache.Client) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: location.New(conn, mc),
		}
	)

	mux.Post("/", h.create)
	mux.Route("/{location_id}", func(mux chi.Router) {
		mux.Get("/", h.get)
		mux.Put("/", h.update)
		mux.Delete("/", h.delete)
	})

	return mux
}
