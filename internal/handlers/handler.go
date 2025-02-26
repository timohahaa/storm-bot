package handlers

import (
	"net/url"
	"strings"
	"time"

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
			link := string([]rune(text)[e.Offset : e.Offset+e.Length])
			if _, err := url.Parse(link); err != nil {
				log.Warnf("not a valid url: %v", link)
				continue
			}
			links = append(links, link)
		}
	}

	return links
}

func parseMonth(text string) uint {
	switch strings.ToLower(text) {
	case "january", "январь":
		return 1
	case "february", "февраль":
		return 2
	case "march", "март":
		return 3
	case "april", "апрель":
		return 4
	case "may", "май":
		return 5
	case "june", "июнь":
		return 6
	case "july", "июль":
		return 7
	case "august", "август":
		return 8
	case "september", "сентябрь":
		return 9
	case "october", "октябрь":
		return 10
	case "november", "ноябрь":
		return 11
	case "december", "декабрь":
		return 12
	}

	// default -> current month
	return uint(time.Now().Month())
}
