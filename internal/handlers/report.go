package handlers

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
)

func (h *Handler) GetReport(c telebot.Context) error {
	var (
		month = parseMonth(c.Message().Payload)
		msg   *telebot.Message
		err   error
	)
	defer func() {
		if err != nil {
			log.Errorf("[bot] (GetReport): %v", err)
		}
	}()

	// 1
	msg, err = c.Bot().Reply(c.Message(), "Starting...")
	if err != nil {
		return c.Reply(fmt.Sprintf("Internal error: %v", err))
	}

	// 2
	if msg, err = c.Bot().Edit(msg, "Querying db..."); err != nil {
		return c.Reply(fmt.Sprintf("Internal error: %v", err))
	}

	ret, err := h.mod.MonthLinkStats(context.Background(), month, h.getUsernameByID)
	if err != nil {
		return c.Reply(fmt.Sprintf("Internal error: %v", err))
	}

	// 3
	if msg, err = c.Bot().Edit(msg, "Building excel file..."); err != nil {
		return c.Reply(fmt.Sprintf("Internal error: %v", err))
	}
	xlsx, err := ret.ToExcel()
	if err != nil {
		return c.Reply(fmt.Sprintf("Internal error: %v", err))
	}

	// 4
	if msg, err = c.Bot().Edit(msg, "Uploading file to Telegram servers..."); err != nil {
		return c.Reply(fmt.Sprintf("Internal error: %v", err))
	}

	if _, err = c.Bot().Edit(msg, &telebot.Document{
		File:     telebot.FromReader(xlsx),
		FileName: fmt.Sprintf("report-%d.xlsx", month),
	}); err != nil {
		return c.Reply(fmt.Sprintf("Internal error: %v", err))
	}

	return nil
}

func (h *Handler) getUsernameByID(id int64) (string, error) {
	chat, err := h.b.ChatByID(id)
	if err != nil {
		return "", err
	}

	return "@" + chat.Username, nil
}
