package store

import (
	"context"

	"portservice/core"
)

var _ core.PortStore = (*PortsMock)(nil)

func NewPortsMock(returnError error) *PortsMock {
	return &PortsMock{ReturnError: returnError}
}

type PortsMock struct {
	ReturnError error
	NewData     map[string]core.Port
}

func (p *PortsMock) GetByID(ctx context.Context, portID string) (core.Port, error) {
	return core.Port{}, nil
}

func (p *PortsMock) BulkInsert(ctx context.Context, ports map[string]core.Port) error {
	p.NewData = ports
	return p.ReturnError
}
