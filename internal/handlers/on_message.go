package handlers

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func (h *Handler) OnMessage(c telebot.Context) error {
	fmt.Printf("%+v\n\n", c.Message().ThreadID)

	c.Send(telebot.ChatID(-1002429016999), "test-send", &telebot.SendOptions{
		ThreadID: 130,
	})
	return nil
}
