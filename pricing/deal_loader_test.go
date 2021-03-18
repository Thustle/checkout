package pricing

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLoadDealsSuccess(t *testing.T) {
	input := "A,3,130\nB,2,45"
	reader := bytes.NewBufferString(input)

	expected := DealStore {
		map[string][]Deal {
			"A": {{"A", 3, 130}},
			"B": {{"B", 2, 45}},
		},
	}
	received, err := LoadDeals(reader)

	if err != nil {
		t.Error("unexpected error", err)
	}
	if !reflect.DeepEqual(expected, *received) {
		t.Errorf("expected %v but received %v", expected, *received)
	}
}

func TestLoadDealsParseErrorInvalidQty(t *testing.T) {
	input := "A,3,130\nB,invalid,45"
	reader := bytes.NewBufferString(input)

	_, err := LoadDeals(reader)

	if err == nil {
		t.Error("expected error is missing")
	}
}

func TestLoadDealsParseErrorInvalidPrice(t *testing.T) {
	input := "A,3,130\nB,2,invalid"
	reader := bytes.NewBufferString(input)

	_, err := LoadDeals(reader)

	if err == nil {
		t.Error("expected error is missing")
	}
}

func TestLoadDealsReadErrorMissingData(t *testing.T) {
	input := "A"
	reader := bytes.NewBufferString(input)

	_, err := LoadDeals(reader)

	if err == nil {
		t.Error("expected error is missing")
	}
}
