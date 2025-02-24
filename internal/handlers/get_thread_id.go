package handlers

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func (h *Handler) GetThreadID(c telebot.Context) error {
	if c.Chat().Type != telebot.ChatSuperGroup {
		c.Reply("Current chat is not a supergroup.")
		return nil
	}

	c.Reply(
		fmt.Sprintf("Current thread id: %d", c.Message().ThreadID),
	)
	return nil
}
