package usecase_test

import (
	"bytes"
	"io"
)

func NewMockFile(filename string, data []byte) *MockFile {
	return &MockFile{
		Filename: filename,
		Data:     bytes.NewReader(data),
	}
}

type MockFile struct {
	Filename string
	Data     io.ReadSeeker
}

func (f *MockFile) Name() string {
	return f.Filename
}

func (f *MockFile) Read(p []byte) (int, error) {
	return f.Data.Read(p)
}

func (f *MockFile) Seek(offset int64, whence int) (int64, error) {
	return f.Data.Seek(offset, whence)
}
