package handlers

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

func (h *Handler) OnMessage(c telebot.Context) error {
	fmt.Printf("%+v\n\n", c.Chat())
	return nil
}
