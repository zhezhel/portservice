package state

import "context"

func NewManagerMock(data map[string]int64) *ManagerMock {
	return &ManagerMock{
		Data:        data,
		ReturnError: nil,
		NewData:     make(map[string]int64),
	}
}

type ManagerMock struct {
	Data        map[string]int64
	ReturnError error
	NewData     map[string]int64
}

func (m *ManagerMock) GetOffset(_ context.Context, filename string) (int64, error) {
	return m.Data[filename], m.ReturnError
}

func (m *ManagerMock) SetOffset(_ context.Context, offset int64, filename string) error {
	m.NewData[filename] = offset
	return nil
}
