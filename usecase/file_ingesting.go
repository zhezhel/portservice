package usecase

import (
	"bufio"
	"context"
	"errors"
	"io"
	"strings"

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
	objectsCountLimit int
	portStore         core.PortStore
	state             stateManager
}

func (f *FileIngestor) Start(ctx context.Context, file File) error {
	offset, err := f.state.GetOffset(file.Name())
	if err != nil {
		return err
	}

	_, err = file.Seek(offset+1, io.SeekStart)
	if err != nil {
		return err
	}

	r := bufio.NewReader(io.MultiReader(strings.NewReader("{"), file))

	dec, err := decoder.NewJSON(r)
	if err != nil {
		return err
	}

	for {
		ports := make(map[string]core.Port)

		for i := 0; i < f.objectsCountLimit; i++ {
			value, err := dec.ReadOne()
			if errors.Is(err, core.ErrPortNotFound) {
				break
			}

			if err != nil {
				break
			}

			ports[value.ID] = *value
		}

		if len(ports) == 0 {
			return nil
		}

		err = f.portStore.BulkInsert(ctx, ports)
		if err != nil {
			return err
		}
	}
}

type stateManager interface {
	GetOffset(filename string) (int64, error)
	SetOffset(offset int64, filename string) error
}

type File interface {
	Name() string
	io.ReadSeeker
}
