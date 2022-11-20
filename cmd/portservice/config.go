package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrFieldRequired = errors.New("field required")

func LoadConfig() (Config, error) {
	var config Config

	config.FileSource = os.Getenv("FILE_SOURCE")
	if config.FileSource == "" {
		return Config{}, fmt.Errorf("%w: FILE_SOURCE", ErrFieldRequired)
	}

	config.PostgresDSN = os.Getenv("PG_DSN")
	if config.PostgresDSN == "" {
		return Config{}, fmt.Errorf("%w: PG_DSN", ErrFieldRequired)
	}

	limitStr := os.Getenv("OBJECTS_COUNT_LIMIT")
	if limitStr == "" {
		return Config{}, fmt.Errorf("%w: OBJECTS_COUNT_LIMIT", ErrFieldRequired)
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return Config{}, fmt.Errorf("invalid value OBJECTS_COUNT_LIMIT: %w", err)
	}
	config.ObjectsCountLimit = limit
	return config, nil
}

type Config struct {
	FileSource        string
	PostgresDSN       string
	ObjectsCountLimit int
}
