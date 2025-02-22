package image

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/timohahaa/gw/internal/modules/image"
)

func New(conn *pgxpool.Pool, mc *memcache.Client, minio *minio.Client) *chi.Mux {
	var (
		mux = chi.NewMux()
		h   = handler{
			mod: image.New(conn, mc, minio),
		}
	)

	mux.Get("/", h.list)
	mux.Post("/upload", h.upload)
	mux.Route("/{image_id}", func(mux chi.Router) {
		mux.Get("/", h.download)
		mux.Delete("/", h.delete)
	})

	return mux
}
