package handlers

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
)

func (h *Handler) OnMessage(c telebot.Context) error {
	links := extractLinks(c.Message())
	if len(links) != 0 {
		if err := h.mod.CreateLinks(
			context.Background(),
			c.Sender().ID,
			c.Chat().ID,
			links,
		); err != nil {
			log.Errorf("[bot] (OnMessage): %v", err)
			return nil
		}
	}
	return nil
}
