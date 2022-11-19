package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"

	"portservice/core"
)

var _ core.PortStore = (*Ports)(nil)

func NewPorts(conn *pgx.Conn) (*Ports, error) {
	return &Ports{
		conn: conn,
	}, nil
}

type Ports struct {
	conn *pgx.Conn
}

func (p *Ports) GetByID(ctx context.Context, portID string) (core.Port, error) {
	return core.Port{}, core.ErrPortNotFound
}

func (p *Ports) BulkInsert(ctx context.Context, ports map[string]core.Port) error {
	return p.conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		query, args := generateBulkInsertSQL(ports)
		_, err := tx.Exec(ctx, query, args...)
		return err
	})
}

func generateBulkInsertSQL(ports map[string]core.Port) (string, []any) {
	q := "INSERT INTO portservice.ports VALUES "
	args := []any{}
	counter := 0
	for id, port := range ports {
		// data, _ := json.Marshal(port)
		args = append(args, id, port)

		q += fmt.Sprintf("($%d, $%d),", 2*counter+1, 2*counter+2)

		counter++
	}

	q = strings.TrimRight(q, ",")

	q += "ON CONFLICT (id) DO UPDATE SET data = EXCLUDED.data"

	return q, args
}
