package decoder_test

import (
	"bytes"
	"io"
)

func NewMockReadSeeker(data []byte) *MockReadSeeker {
	return &MockReadSeeker{
		Data: bytes.NewReader(data),
	}
}

type MockReadSeeker struct {
	ReadCalled int64
	SeekOffset int64
	SeekWhence int
	Data       io.ReadSeeker
}

func (f *MockReadSeeker) Read(p []byte) (int, error) {
	f.ReadCalled++
	return f.Data.Read(p)
}

func (f *MockReadSeeker) Seek(offset int64, whence int) (int64, error) {
	f.SeekOffset = offset
	f.SeekWhence = whence
	return f.Data.Seek(offset, whence)
}
