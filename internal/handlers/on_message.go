package handlers

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/react"
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

		// note to user that their message was processed
		c.Bot().React(c.Recipient(), c.Message(), telebot.Reactions{
			Reactions: []telebot.Reaction{react.ThumbUp}},
		)
	}
	return nil
}
