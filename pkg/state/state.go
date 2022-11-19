package state

import "github.com/jackc/pgx/v4"

func NewStateManager(conn *pgx.Conn) (*StateManager, error) {
	return &StateManager{
		conn: conn,
	}, nil
}

type StateManager struct {
	conn *pgx.Conn
}

func (s *StateManager) GetOffset(filename string) (int64, error) {
	return 0, nil
}

func (s *StateManager) SetOffset(offset int64, filename string) error {
	return nil
}
