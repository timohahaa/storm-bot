package auth

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/gw/config"
	"github.com/timohahaa/gw/internal/modules/customer"
	"github.com/timohahaa/gw/internal/modules/user"
	"github.com/timohahaa/gw/internal/services/gmail"
	"github.com/timohahaa/gw/internal/services/p1sms"
)

func New(
	conn *pgxpool.Pool,
	mc *memcache.Client,
	p1smsApi *p1sms.API,
	gmailSMTP *gmail.Sender,
	cfg *config.Config,
) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: mod{
				customer: customer.New(conn, mc, p1smsApi),
				user:     user.New(conn, mc, gmailSMTP),
			},
			customerAuthCookieName: cfg.CustomerAuthCookieName,
			userAuthCookieName:     cfg.UserAuthCookieName,
			cookieDomain:           cfg.DomainName,
			cookieSecure:           cfg.DomainSecure,
		}
	)

	mux.Route("/customers", func(mux chi.Router) {
		mux.Post("/register", h.createCustomer)
		mux.Post("/login", h.loginCustomer)

		mux.Route("/{user_id}", func(mux chi.Router) {
			mux.Post("/confirm", h.confirmCustomer)
			mux.Post("/send-code", h.sendCodeCustomer)
		})
	})

	mux.Route("/users", func(mux chi.Router) {
		mux.Post("/register", h.createUser)
		mux.Post("/login", h.loginUser)

		mux.Route("/{user_id}", func(mux chi.Router) {
			mux.Post("/confirm", h.confirmUser)
			mux.Post("/send-code", h.sendCodeUser)
		})
	})

	return mux
}
