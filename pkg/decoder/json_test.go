package decoder_test

import (
	"errors"
	"testing"

	"portservice/core"
	"portservice/pkg/decoder"
)

func TestJsonDecoderSkipFirstCharacter(t *testing.T) {
	expectedOffset := int64(1)
	r := NewMockReadSeeker([]byte(""))

	_, _ = decoder.NewJSON(r, 0)

	if r.SeekOffset != expectedOffset {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", r.SeekOffset, expectedOffset)
	}
}

func TestJsonDecoderReturnErrorOnEOF(t *testing.T) {
	r := NewMockReadSeeker([]byte(""))
	dec, _ := decoder.NewJSON(r, 0)

	_, err := dec.ReadOne()

	if !errors.Is(err, core.ErrNoMorePortElements) {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", err, core.ErrNoMorePortElements)
	}
}
