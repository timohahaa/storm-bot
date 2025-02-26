package handlers

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
)

func (h *Handler) GetReport(c telebot.Context) error {
	month := parseMonth(c.Message().Payload)
	ret, err := h.mod.MonthLinkStats(context.Background(), month, h.getUsernameByID)
	if err != nil {
		return nil
	}

	xlsx, err := ret.ToExcel()
	if err != nil {
		log.Errorf("[bot] (GetReport): %v", err)
		c.Reply(fmt.Sprintf("Internal error: %v", err))
		return err
	}

	_, err = h.b.Reply(c.Message(), &telebot.Document{
		File:     telebot.FromReader(xlsx),
		FileName: fmt.Sprintf("report-%d.xlsx", month),
	})

	if err != nil {
		log.Errorf("[bot] (GetReport): %v", err)
		c.Reply(fmt.Sprintf("Internal error: %v", err))
		return err
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
