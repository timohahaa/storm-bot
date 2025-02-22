package business

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/timohahaa/gw/internal/handlers/http/v1/business/analytics"
	"github.com/timohahaa/gw/internal/handlers/http/v1/business/image"
	"github.com/timohahaa/gw/internal/handlers/http/v1/business/link"
	"github.com/timohahaa/gw/internal/handlers/http/v1/business/location"
	"github.com/timohahaa/gw/internal/handlers/http/v1/business/pcard"
	"github.com/timohahaa/gw/internal/middleware"
	"github.com/timohahaa/gw/internal/modules/business"
)

func New(conn *pgxpool.Pool, mc *memcache.Client, minio *minio.Client) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: business.New(conn, mc),
		}
	)

	mux.Post("/", h.create)
	mux.Route("/{business_id}", func(mux chi.Router) {
		mux.Use(middleware.CheckBusinessAccess(conn))

		mux.Get("/", h.get)
		mux.Put("/", h.update)
		mux.Delete("/", h.delete)

		mux.Mount("/punch-cards", pcard.New(conn, mc))
		mux.Mount("/locations", location.New(conn, mc))
		mux.Mount("/images", image.New(conn, mc, minio))
		mux.Mount("/analytics", analytics.New(conn))
		mux.Mount("/links", link.New(conn))
		// TODO: maybe create a separate handler for adding members to a business
	})

	return mux
}
