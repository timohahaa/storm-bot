package customer

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/internal/handlers/http/v1/customer/card"
	"github.com/timohahaa/gw/internal/modules/customer"
	"github.com/timohahaa/gw/internal/services/p1sms"
)

func New(conn *pgxpool.Pool, mc *memcache.Client, p1smsApi *p1sms.API) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: customer.New(conn, mc, p1smsApi),
		}
	)

	mux.Route("/{customer_id}", func(mux chi.Router) {
		mux.Get("/", h.get)
		mux.Put("/", h.update)

		mux.Mount("/cards", card.New(conn, mc))
	})

	return mux
}
