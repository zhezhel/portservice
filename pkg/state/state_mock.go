package state

func NewStateManagerMock(data map[string]int64) *StateManagerMock {
	return &StateManagerMock{}
}

type StateManagerMock struct {
	Data         map[string]int64
	Return_error error
	NewData      map[string]int64
}

func (m *StateManagerMock) GetOffset(filename string) (int64, error) {
	return m.Data[filename], m.Return_error
}

func (m *StateManagerMock) SetOffset(offset int64, filename string) error {
	m.NewData[filename] = offset
	return nil
}
