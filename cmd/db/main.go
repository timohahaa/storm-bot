package main

import (
	"errors"
	"flag"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	attempts = 20
	timeout  = time.Second
	dsn      = flag.String("dsn", "", "postgres dsn")
)

func main() {
	flag.Parse()
	var m *migrate.Migrate
	var err error

	for attempts > 0 {
		m, err = migrate.New("file://migrations", *dsn)
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

	err = m.Down()
	defer func() {
		_, _ = m.Close()
	}()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("[migrate]: migrating down error: %s", err)
	}
	if errors.Is(err, migrate.ErrNoChange) {
		log.Infof("[migrate]: no change.")
		return
	}

	log.Infof("[migrate]: migration down: success.")
}
