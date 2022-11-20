package usecase_test

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"portservice/core"
	"portservice/pkg/state"
	"portservice/pkg/store"
	"portservice/usecase"
)

func TestIngesting_FromStart(t *testing.T) {
	stateManager := state.NewManagerMock(map[string]int64{})
	portStore := store.NewPortsMock()
	ingestor, _ := usecase.NewFileIngestor(portStore, stateManager, 100)
	data, _ := portsToJSON(portAjman)
	file := NewMockFile("test.json", data)

	_ = ingestor.Start(context.Background(), file)

	actual := portStore.NewData[portAjman.ID]
	if !reflect.DeepEqual(actual, portAjman) {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", actual, portAjman)
	}
}

func TestIngesting_WithOffset(t *testing.T) {
	stateManager := state.NewManagerMock(map[string]int64{"test.json": 207})
	portStore := store.NewPortsMock()
	ingestor, _ := usecase.NewFileIngestor(portStore, stateManager, 100)
	data, _ := portsToJSON(portAjman, portAbuDhabi)
	file := NewMockFile("test.json", data)

	_ = ingestor.Start(context.Background(), file)

	actual := portStore.NewData[portAbuDhabi.ID]
	if !reflect.DeepEqual(actual, portAbuDhabi) {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", actual, portAjman)
	}
}

func TestIngesting_WithIncorrectOffset(t *testing.T) {
	incorrectOffset := int64(192)
	stateManager := state.NewManagerMock(map[string]int64{"test.json": incorrectOffset})
	portStore := store.NewPortsMock()
	ingestor, _ := usecase.NewFileIngestor(portStore, stateManager, 100)
	data, _ := portsToJSON(portAjman, portAbuDhabi)
	file := NewMockFile("test.json", data)

	err := ingestor.Start(context.Background(), file)

	if err == nil {
		t.Fatalf("\nwant an error, got nil")
	}
}

func TestIngesting_EndOfFile(t *testing.T) {
	stateManager := state.NewManagerMock(map[string]int64{"test.json": 412})
	portStore := store.NewPortsMock()
	ingestor, _ := usecase.NewFileIngestor(portStore, stateManager, 100)
	data, _ := portsToJSON(portAjman, portAbuDhabi)
	file := NewMockFile("test.json", data)

	_ = ingestor.Start(context.Background(), file)

	actual := len(portStore.NewData)
	if actual != 0 {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", actual, 0)
	}
}

func TestIngesting_TwoPortsWithSameID(t *testing.T) {
	stateManager := state.NewManagerMock(map[string]int64{})
	portStore := store.NewPortsMock()
	ingestor, _ := usecase.NewFileIngestor(portStore, stateManager, 1)
	expected := portAjman
	expected.City = "Warsaw"
	expected.Code = "123"
	expected.Name = "Test name"
	data, _ := portsToJSON(portAjman, expected)
	file := NewMockFile("test.json", data)

	_ = ingestor.Start(context.Background(), file)

	actual := portStore.NewData[expected.ID]
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", actual, expected)
	}
}

func portsToJSON(ports ...core.Port) ([]byte, error) {
	tmp := make(map[string]core.Port)

	for i := range ports {
		tmp[ports[i].ID] = ports[i]
	}

	return json.Marshal(tmp)
}

var portAjman = core.Port{
	ID:          "AEAJM",
	Name:        "Ajman",
	Coordinates: []float64{55.5136433, 25.4052165},
	City:        "Ajman",
	Country:     "United Arab Emirates",
	Alias:       []string{},
	Regions:     []string{},
	Timezone:    "Asia/Dubai",
	Unlocs:      []string{"AEAJM"},
	Code:        "52000",
}

var portAbuDhabi = core.Port{
	ID:          "AEAUH",
	Name:        "Abu Dhabi",
	Coordinates: []float64{54.37, 24.47},
	City:        "Abu Dhabi",
	Country:     "United Arab Emirates",
	Alias:       []string{},
	Regions:     []string{},
	Timezone:    "Asia/Dubai",
	Unlocs:      []string{"AEAUH"},
	Code:        "52001",
}
