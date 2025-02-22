package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/storm-bot/config"
	"github.com/timohahaa/storm-bot/internal/bot"
	"gopkg.in/telebot.v4"
)

type Handler struct {
	mod *bot.Module
	cfg config.Config
	b   *telebot.Bot
}

func New(b *telebot.Bot, conn *pgxpool.Pool, cfg config.Config) *Handler {
	return &Handler{
		mod: bot.New(conn),
		cfg: cfg,
		b:   b,
	}
}
