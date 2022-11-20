package decoder

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"portservice/core"
)

var (
	ErrUnexpectedType = errors.New("unexpected type")
)

func NewJSON(r io.ReadSeeker, offset int64) (*JSON, error) {
	_, err := r.Seek(offset+1, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// "{" added to put the decoder in the correct state after seek
	reader := bufio.NewReader(io.MultiReader(strings.NewReader("{"), r))
	decoder := json.NewDecoder(reader)
	_, err = decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("decoder error: %w", err)
	}

	return &JSON{
		decoder: *decoder,
		offset:  offset,
	}, nil
}

type JSON struct {
	decoder json.Decoder
	offset  int64
}

func (j *JSON) ReadOne() (core.Port, error) {
	if !j.decoder.More() {
		return core.Port{}, core.ErrNoMorePortElements
	}
	t, err := j.decoder.Token()
	if err != nil {
		return core.Port{}, fmt.Errorf("decoder error: %w", err)
	}

	if t == json.Delim('}') {
		return core.Port{}, fmt.Errorf("decoder error: %w", core.ErrNoMorePortElements)
	}

	id, ok := t.(string)
	if !ok {
		return core.Port{}, fmt.Errorf("decoder error: %w", ErrUnexpectedType)
	}

	var value core.Port
	if err := j.decoder.Decode(&value); err != nil {
		return core.Port{}, fmt.Errorf("decoder error: %w", err)
	}

	value.ID = id

	return value, nil
}

func (j *JSON) InputOffset() int64 {
	return j.offset + j.decoder.InputOffset()
}
