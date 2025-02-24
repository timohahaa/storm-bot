package service

import (
	"github.com/timohahaa/storm-bot/internal/handlers"
	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func (srv *Service) route() {
	var (
		h = handlers.New(srv.bot, srv.conn, *srv.cfg)
	)

	srv.bot.Handle(telebot.OnText, h.OnMessage)
	srv.bot.Handle("/threadid", h.GetThreadID)

	adminOnly := srv.bot.Group()
	adminOnly.Use(middleware.Whitelist(srv.cfg.AdminIDs...))
	adminOnly.Handle("/threadid", h.GetThreadID)
}
