package v1

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/timohahaa/gw/config"
	"github.com/timohahaa/gw/internal/handlers/http/v1/auth"
	"github.com/timohahaa/gw/internal/handlers/http/v1/business"
	"github.com/timohahaa/gw/internal/handlers/http/v1/customer"
	"github.com/timohahaa/gw/internal/handlers/http/v1/user"
	"github.com/timohahaa/gw/internal/middleware"
	"github.com/timohahaa/gw/internal/services/gmail"
	"github.com/timohahaa/gw/internal/services/p1sms"
)

func New(
	conn *pgxpool.Pool,
	mc *memcache.Client,
	minio *minio.Client,
	p1smsApi *p1sms.API,
	gmailSMTP *gmail.Sender,
	cfg *config.Config,
) *chi.Mux {
	var (
		mux = chi.NewMux()
	)

	mux.Mount("/auth", auth.New(conn, mc, p1smsApi, gmailSMTP, cfg))
	mux.Group(func(mux chi.Router) {
		mux.Use(middleware.UserAuth(conn, cfg.UserAuthCookieName))

		mux.Mount("/businesses", business.New(conn, mc, minio))
		mux.Mount("/users", user.New(conn, mc, gmailSMTP))
	})

	mux.Group(func(mux chi.Router) {
		mux.Use(middleware.CustomerAuth(conn, cfg.CustomerAuthCookieName))

		mux.Mount("/customers", customer.New(conn, mc, p1smsApi))
	})

	return mux
}
