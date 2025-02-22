package pcard

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/internal/modules/card"
	"github.com/timohahaa/gw/internal/modules/pcard"
)

func New(conn *pgxpool.Pool, mc *memcache.Client) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: mod{
				pcard: pcard.New(conn, mc),
				card:  card.New(conn, mc),
			},
		}
	)

	mux.Get("/", h.list)
	mux.Post("/", h.create)
	mux.Route("/{punch_card_id}", func(mux chi.Router) {
		mux.Get("/", h.get)
		mux.Put("/", h.update)
		mux.Delete("/", h.delete)

		mux.Post("/stamp", h.stamp)
	})

	return mux
}
