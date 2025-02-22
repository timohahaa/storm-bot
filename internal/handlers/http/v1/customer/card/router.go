package card

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/internal/modules/card"
)

func New(conn *pgxpool.Pool, mc *memcache.Client) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: card.New(conn, mc),
		}
	)

	mux.Get("/", h.list)
	mux.Post("/", h.create)
	mux.Route("/{card_id}", func(mux chi.Router) {
		mux.Get("/", h.get)
		mux.Delete("/", h.delete)
	})

	return mux
}
