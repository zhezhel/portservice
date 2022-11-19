package core

import (
	"context"
	"errors"
)

var ErrPortNotFound = errors.New("port not found")

type Port struct {
	ID          string    `json:"-"`
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
}

type PortStore interface {
	GetByID(ctx context.Context, portID string) (Port, error)
	BulkInsert(ctx context.Context, ports map[string]Port) error
}
