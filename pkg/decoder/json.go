package decoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"portservice/core"
)

func NewJSON(r io.Reader) (*JSON, error) {
	decoder := json.NewDecoder(r)
	t, err := decoder.Token()
	if err != nil {
		return nil, err
	}
	if t != json.Delim('{') {
		return nil, fmt.Errorf("unexpected token: %s", t)
	}

	return &JSON{
		decoder: *decoder,
	}, nil
}

type JSON struct {
	decoder json.Decoder
}

func (j *JSON) ReadOne() (*core.Port, error) {
	if !j.decoder.More() {
		return nil, errors.New("no more elements to read")
	}
	t, err := j.decoder.Token()
	if err != nil {
		return nil, err
	}

	if t == json.Delim('}') {
		return nil, errors.New("no more elements to read")
	}

	id, ok := t.(string)
	if !ok {
		return nil, errors.New("unknown type")
	}

	var value core.Port
	if err := j.decoder.Decode(&value); err != nil {
		return nil, err
	}

	value.ID = id

	return &value, nil
}

func (j *JSON) InputOffset() int64 {
	return j.decoder.InputOffset()
}
