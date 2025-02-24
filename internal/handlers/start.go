package handlers

import (
	"gopkg.in/telebot.v4"
)

func (h *Handler) OnStart(c telebot.Context) error {
	if c.Chat().Type == telebot.ChatPrivate {
		c.Reply("Hi there! This bot is only usefull inside a supergroup or a group-chat.")
	}
	return nil
}
