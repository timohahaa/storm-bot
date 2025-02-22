package service

import (
	"github.com/timohahaa/storm-bot/internal/handlers"
	"gopkg.in/telebot.v4"
)

func (srv *Service) route() {
	var (
		h = handlers.New(srv.conn, *srv.cfg)
	)

	srv.bot.Handle(telebot.OnText, h.OnMessage)
}
