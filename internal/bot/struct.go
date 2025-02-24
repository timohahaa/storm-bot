package bot

import (
	"github.com/google/uuid"
)

type Link struct {
	ID     uuid.UUID `db:"id"      json:"id"`
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	ChatID int64     `db:"chat_id" json:"chat_id"`
	Link   string    `db:"link"    json:"link"`
}

type User struct {
	ID         uuid.UUID `db:"id"          json:"id"`
	TelegramID int64     `db:"telegram_id" json:"telegram_id"`
	IsAdmin    bool      `db:"is_admin"    json:"is_admin"`
}

type UserLink struct {
	UserID int64  `db:"user_id" json:"user_id"`
	Link   string `db:"link"    json:"link"`
}
