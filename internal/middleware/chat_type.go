package middleware

import "gopkg.in/telebot.v4"

func IsGroup() telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			switch c.Chat().Type {
			case telebot.ChatGroup, telebot.ChatSuperGroup:
				return next(c)
			case telebot.ChatPrivate:
				c.Reply("This bot is for use in groups only!")
				return nil
			}
			// telebot.ChatChannel
			//telebot.ChatChannelPrivate
			return nil
		}
	}
}
