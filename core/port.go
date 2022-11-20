package core

import (
	"context"
	"errors"
)

var ErrPortNotFound = errors.New("port not found")
var ErrNoMorePortElements = errors.New("no more port elements to read")

type Port struct {
	ID          string    `json:"-"`
	Name        string    `json:"name"`
	Coordinates []float64 `json:"coordinates"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type PortStore interface {
	BulkInsert(ctx context.Context, ports map[string]Port) error
}
