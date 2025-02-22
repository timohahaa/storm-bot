package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/timohahaa/storm-bot/config"
	"gopkg.in/telebot.v4"
)

type Service struct {
	cfg *config.Config

	conn *pgxpool.Pool
	bot  *telebot.Bot

	signal chan os.Signal
}

func New(cfg *config.Config) (*Service, error) {
	var (
		s = &Service{
			cfg:    cfg,
			signal: make(chan os.Signal),
		}
		err error
	)
	s.runMigrations()

	if s.conn, err = pgxpool.New(context.Background(), cfg.PostgresDSN); err != nil {
		log.Errorf("[service] error connecting to postgres: %v", err)
		return nil, err
	}

	s.route()

	return s, nil
}

func (srv *Service) Run() error {
	var (
		signals = []os.Signal{
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGKILL,
		}
	)

	signal.Notify(srv.signal, signals...)
	signal := <-srv.signal

	log.Infof("[service] handled signal: %v", signal)
	log.Info("[service] shutting down...")

	return nil
}
