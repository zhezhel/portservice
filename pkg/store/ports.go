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

func (p *Ports) BulkInsert(ctx context.Context, ports map[string]core.Port) error {
	return p.conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		query, args := generateBulkInsertSQL(ports)
		_, err := tx.Exec(ctx, query, args...)
		return err
	})
}

func generateBulkInsertSQL(ports map[string]core.Port) (query string, args []any) {
	query = "INSERT INTO portservice.ports VALUES "
	counter := 0
	for id := range ports {
		args = append(args, id, ports[id])

		query += fmt.Sprintf("($%d, $%d),", 2*counter+1, 2*counter+2)

		counter++
	}

	query = strings.TrimRight(query, ",")

	query += " ON CONFLICT (id) DO UPDATE SET data = EXCLUDED.data"

	return query, args
}
