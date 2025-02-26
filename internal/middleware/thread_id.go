package middleware

import "gopkg.in/telebot.v4"

func ThreadID(id int) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			if c.Message() != nil && c.Message().ThreadID == id {
				next(c)
				return nil
			}
			return nil
		}
	}
}
