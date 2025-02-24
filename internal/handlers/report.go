package handlers

import (
	"context"
	"fmt"

	"gopkg.in/telebot.v4"
)

func (h *Handler) GetReport(c telebot.Context) error {
	month := parseMonth(c.Message().Payload)
	ret, err := h.mod.MonthLinkStats(context.Background(), month)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(ret)
	return nil
}
