package service

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/timohahaa/gw/docs"
	v1 "github.com/timohahaa/gw/internal/handlers/http/v1"
	"github.com/timohahaa/gw/internal/services/gmail"
	"github.com/timohahaa/gw/internal/services/p1sms"
)

func (srv *Service) route() {
	var (
		p1smsApi  = p1sms.New(srv.cfg.P1SMSApiKey)
		gmailSMTP = gmail.New(
			srv.cfg.SMTPHostname,
			srv.cfg.SMTPPort,
			srv.cfg.SMTPUsername,
			srv.cfg.SMTPPassword,
		)
	)

	srv.mux.Mount("/api/v1", v1.New(srv.conn, srv.mc, srv.minio, p1smsApi, gmailSMTP, srv.cfg))
	srv.mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	srv.mux.Mount("/swagger", httpSwagger.WrapHandler)
}
