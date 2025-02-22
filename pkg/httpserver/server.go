package httpserver

import (
	"context"
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	server          *http.Server
	shutDownTimeout time.Duration
}

func New(bindAddr string, handler http.Handler, opts ...Option) *Server {
	stdServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		Addr:         bindAddr,
	}

	s := &Server{
		server:          stdServer,
		shutDownTimeout: defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	log.Infof("[HTTP server] listening on: %s", s.server.Addr)
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Warn("[HTTP server] closed.")
			} else {
				log.Fatalf("[HTTP server] error: %v", err)
			}
		}
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutDownTimeout)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		log.Errorf("[HTTP server] shutdown error: %v", err)
		return err
	}

	return nil
}
