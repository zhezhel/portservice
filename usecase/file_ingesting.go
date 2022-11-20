package usecase

import (
	"context"
	"errors"
	"io"

	"portservice/core"
	"portservice/pkg/decoder"
)

func NewFileIngestor(portStore core.PortStore, stateManager stateManager, objectsCountLimit int) (*FileIngestor, error) {
	return &FileIngestor{
		portStore:         portStore,
		state:             stateManager,
		objectsCountLimit: objectsCountLimit,
	}, nil
}

type FileIngestor struct {
	portStore         core.PortStore
	state             stateManager
	objectsCountLimit int
}

func (f *FileIngestor) Start(ctx context.Context, file File) error {
	offset, err := f.state.GetOffset(ctx, file.Name())
	if err != nil {
		return err
	}

	dec, err := decoder.NewJSON(file, offset)
	if err != nil {
		return err
	}

	for {
		ports := make(map[string]core.Port)

		for i := 0; i < f.objectsCountLimit; i++ {
			value, err := dec.ReadOne()
			if errors.Is(err, core.ErrNoMorePortElements) {
				break
			}
			if err != nil {
				return err
			}

			ports[value.ID] = value
		}

		if len(ports) == 0 {
			return nil
		}

		err = f.portStore.BulkInsert(ctx, ports)
		if err != nil {
			return err
		}

		err = f.state.SetOffset(ctx, dec.InputOffset(), file.Name())
		if err != nil {
			return err
		}
	}
}

type stateManager interface {
	GetOffset(ctx context.Context, filename string) (int64, error)
	SetOffset(ctx context.Context, offset int64, filename string) error
}

type File interface {
	Name() string
	io.ReadSeeker
}
