package state

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func NewManager(conn *pgx.Conn) (*Manager, error) {
	return &Manager{
		conn: conn,
	}, nil
}

type Manager struct {
	conn *pgx.Conn
}

func (s *Manager) GetOffset(ctx context.Context, filename string) (int64, error) {
	var offset int64
	err := s.conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		query := "SELECT COALESCE((SELECT offset_  FROM portservice.states WHERE filename = $1), 0)"
		args := []any{filename}
		row := tx.QueryRow(ctx, query, args...)
		return row.Scan(&offset)
	})

	return offset, err
}

func (s *Manager) SetOffset(ctx context.Context, offset int64, filename string) error {
	return s.conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		query := "INSERT INTO portservice.states (filename, offset_) VALUES ($1, $2) "
		query += "ON CONFLICT (filename) DO UPDATE SET offset_ = EXCLUDED.offset_"
		args := []any{filename, offset}
		_, err := tx.Exec(ctx, query, args...)
		return err
	})
}
