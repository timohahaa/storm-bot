package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	attempts = 20
	timeout  = time.Second
)

func (srv *Service) runMigrations() {
	dbURL := srv.cfg.PostgresDSN
	dbURL += "?sslmode=disable"

	var m *migrate.Migrate
	var err error

	for attempts > 0 {
		m, err = migrate.New("file://migrations", dbURL)
		if err == nil {
			break
		}
		log.Infof("[migrate]: trying to connect, attempts left: %d", attempts)
		time.Sleep(timeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("[migrate]: connect error: %s", err)
	}

	err = m.Up()
	defer func() {
		_, _ = m.Close()
	}()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("[migrate]: migrating up error: %s", err)
	}
	if errors.Is(err, migrate.ErrNoChange) {
		log.Infof("[migrate]: no change.")
		return
	}

	log.Infof("[migrate]: migration up: success.")
}
