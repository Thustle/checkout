package pricing

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLoadPricesSuccess(t *testing.T) {
	input := "A,50\nB,30\nC,20\nD,15"
	reader := bytes.NewBufferString(input)

	expected := map[string]int {
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
	received, err := LoadPrices(reader)

	if err != nil {
		t.Error("unexpected error", err)
	}
	if !reflect.DeepEqual(expected, *received) {
		t.Errorf("expected %v but received %v", expected, *received)
	}
}

func TestLoadPricesParseErrorInvalidPrice(t *testing.T) {
	input := "A,50\nB,30\nC,invalid\nD,15"
	reader := bytes.NewBufferString(input)

	_, err := LoadPrices(reader)

	if err == nil {
		t.Error("expected error is missing")
	}
}

func TestLoadPricesReadErrorMissingData(t *testing.T) {
	input := "A"
	reader := bytes.NewBufferString(input)

	_, err := LoadPrices(reader)

	if err == nil {
		t.Error("expected error is missing")
	}
}
