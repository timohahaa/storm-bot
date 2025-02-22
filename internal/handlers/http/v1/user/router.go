package user

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/internal/modules/user"
	"github.com/timohahaa/gw/internal/services/gmail"
)

func New(conn *pgxpool.Pool, mc *memcache.Client, gmailSMTP *gmail.Sender) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: user.New(conn, mc, gmailSMTP),
		}
	)

	mux.Route("/{user_id}", func(mux chi.Router) {
		mux.Get("/", h.get)
		mux.Put("/", h.update)
	})

	return mux
}
