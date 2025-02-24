package service

import (
	"github.com/timohahaa/storm-bot/internal/handlers"
	"github.com/timohahaa/storm-bot/internal/middleware"
	"gopkg.in/telebot.v4"
	mw "gopkg.in/telebot.v4/middleware"
)

func (srv *Service) route() {
	var (
		h = handlers.New(srv.bot, srv.conn, *srv.cfg)
	)

	srv.bot.Handle(telebot.OnText, h.OnMessage, middleware.IsGroup())
	srv.bot.Handle("/start", h.OnStart)

	adminOnly := srv.bot.Group()
	adminOnly.Use(mw.Whitelist(srv.cfg.AdminIDs...))
	adminOnly.Handle("/threadid", h.GetThreadID)
	adminOnly.Handle("/report", h.GetReport)
}
