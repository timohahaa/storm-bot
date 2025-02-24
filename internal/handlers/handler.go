package handlers

import (
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
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

func extractLinks(msg *telebot.Message) []string {
	var (
		text  = msg.Text
		links []string
	)

	for _, e := range msg.Entities {
		switch e.Type {
		case telebot.EntityTextLink:
			links = append(links, e.URL)
		case telebot.EntityURL:
			link := text[e.Offset:e.Length]
			if _, err := url.Parse(link); err != nil {
				log.Warnf("not a valid url: %v", link)
				continue
			}
			links = append(links, link)
		}
	}

	return links
}
