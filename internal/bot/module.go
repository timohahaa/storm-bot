package bot

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *Module {
	return &Module{
		conn: conn,
	}
}

//func (m *Module) CreateUser(ctx context.Context, telegramID int64, isAdmin bool) (User, error) {
//	row := m.conn.QueryRow(ctx, createUserQuery, telegramID, isAdmin)
//	var u User
//	err := row.Scan(
//		&u.ID,
//		&u.TelegramID,
//		&u.IsAdmin,
//	)
//	return u, err
//}
//
//func (m *Module) GetUser(ctx context.Context, telegramID int64) (User, error) {
//	row := m.conn.QueryRow(ctx, getUserQuery, telegramID)
//	var u User
//	err := row.Scan(
//		&u.ID,
//		&u.TelegramID,
//		&u.IsAdmin,
//	)
//	return u, err
//}

//func (m *Module) CreateLink(ctx context.Context, userID, chatID int64, link string) (Link, error) {
//	row := m.conn.QueryRow(ctx, createLinkQuery, userID, chatID, link)
//	var l Link
//	err := row.Scan(
//		&l.ID,
//		&l.UserID,
//		&l.ChatID,
//		&l.Link,
//	)
//	return l, err
//}

func (m *Module) CreateLinks(ctx context.Context, userID, chatID int64, links []string) error {
	tx, err := m.conn.Begin(ctx)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(
		ctx,
		"create-links",
		createLinkQueryNoReturning,
	)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	for _, link := range links {
		if _, err := tx.Exec(ctx, stmt.Name, userID, chatID, link); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

func (m *Module) MonthLinkStats(
	ctx context.Context,
	month uint,
	getUsernameFunc func(int64) (string, error),
) (map[string][]string, error) {
	rows, err := m.conn.Query(ctx, monthLinkStatsQuery, month)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var stats = make(map[string][]string)
	for rows.Next() {
		var i UserLink
		if err := rows.Scan(
			&i.UserID,
			&i.Link,
		); err != nil {
			return nil, err
		}

		username, err := getUsernameFunc(i.UserID)
		if err != nil {
			return nil, err
		}

		if links, ok := stats[username]; ok {
			stats[username] = append(links, i.Link)
		} else {
			stats[username] = []string{i.Link}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return stats, nil
}
