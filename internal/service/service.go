package service

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"github.com/timohahaa/gw/config"
	"github.com/timohahaa/gw/pkg/httpserver"
)

type Service struct {
	cfg *config.Config

	mux   *chi.Mux
	conn  *pgxpool.Pool
	mc    *memcache.Client
	minio *minio.Client

	signal chan os.Signal
}

func New(cfg *config.Config) (*Service, error) {
	var (
		s = &Service{
			cfg:    cfg,
			mux:    chi.NewMux(),
			signal: make(chan os.Signal),
		}
		err error
	)

	if s.conn, err = pgxpool.New(context.Background(), cfg.PostgresDSN); err != nil {
		log.Errorf("[service] error connecting to postgres: %v", err)
		return nil, err
	}

	s.mc = memcache.New(cfg.MemcachedAddrs...)
	if err := s.mc.Ping(); err != nil {
		log.Warnf("[service] memcached ping err: %v", err)
	}

	if s.minio, err = minio.New(cfg.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Secure: cfg.S3UseSSL,
	}); err != nil {
		log.Errorf("[service] error connecting to minio: %v", err)
		return nil, err
	}

	s.route()

	return s, nil
}

func (srv *Service) Run() error {
	var (
		server = httpserver.New(
			net.JoinHostPort("", srv.cfg.HttpPort),
			srv.mux,
		)
		signals = []os.Signal{
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGKILL,
		}
	)

	server.Start()

	signal.Notify(srv.signal, signals...)
	signal := <-srv.signal

	log.Infof("[service] handled signal: %v", signal)
	log.Info("[service] shutting down...")

	_ = server.Shutdown()

	return nil
}
