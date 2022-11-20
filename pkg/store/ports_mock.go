package store

import (
	"context"

	"portservice/core"
)

var _ core.PortStore = (*PortsMock)(nil)

func NewPortsMock() *PortsMock {
	return &PortsMock{}
}

type PortsMock struct {
	ReturnError error
	NewData     map[string]core.Port
}

func (p *PortsMock) BulkInsert(_ context.Context, ports map[string]core.Port) error {
	p.NewData = ports
	return p.ReturnError
}
