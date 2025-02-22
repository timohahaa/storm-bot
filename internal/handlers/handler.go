package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timohahaa/storm-bot/config"
	"github.com/timohahaa/storm-bot/internal/bot"
)

type Handler struct {
	mod *bot.Module
	cfg config.Config
}

func New(conn *pgxpool.Pool, cfg config.Config) *Handler {
	return &Handler{
		mod: bot.New(conn),
		cfg: cfg,
	}
}
