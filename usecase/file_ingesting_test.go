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
}

func TestIngesting_FromStart(t *testing.T) {
	// Arrange
	stateManager := state.NewStateManagerMock(map[string]int64{})
	portStore := store.NewPortsMock(nil)
	ingestor, _ := usecase.NewFileIngestor(portStore, stateManager, 10)
	data, _ := portsToJson(portAjman)
	file := NewMockFile("test.json", data)
	// Act

	_ = ingestor.Start(context.Background(), file)

	// Assert
	actual := portStore.NewData[portAjman.ID]

	if !reflect.DeepEqual(actual, portAjman) {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", actual, portAjman)
	}
}

func TestIngesting_WithOffset(t *testing.T) {

}

func TestIngesting_EndOfFile(t *testing.T) {

}

func portsToJson(ports ...core.Port) ([]byte, error) {
	tmp := make(map[string]core.Port)

	for _, port := range ports {
		tmp[port.ID] = port
	}

	return json.Marshal(tmp)
}
