package store

import (
	"reflect"
	"testing"

	"portservice/core"
)

func TestGenerateBulkInsertSQL_ReturnCorrectQuery(t *testing.T) {
	expected := "INSERT INTO portservice.ports VALUES ($1, $2) "
	expected += "ON CONFLICT (id) DO UPDATE SET data = EXCLUDED.data"
	ports := map[string]core.Port{"first": {Name: "Warsaw"}}

	actual, _ := generateBulkInsertSQL(ports)

	if actual != expected {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", actual, expected)
	}
}

func TestGenerateBulkInsertSQL_ReturnCorrectArgs(t *testing.T) {
	port := core.Port{Name: "Warsaw"}
	id := "first"
	ports := map[string]core.Port{id: port}
	expected := []any{id, port}

	_, actual := generateBulkInsertSQL(ports)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("\nhave: %+v\nwant: %+v\n", actual, expected)
	}
}
