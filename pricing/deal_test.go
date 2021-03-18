package pricing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateDealStoreSuccess(t *testing.T) {
	testDeals := map[string][]Deal{
		"A": {
			{"A", 5, 45},
			{"A", 10, 80},
		},
		"B": {
			{ "B", 2, 45},
		},
	}

	expected := DealStore {
		deals: map[string][]Deal {
			"A": {{"A", 10, 80}, {"A", 5, 45}},
			"B": {{"B", 2, 45}},
		},
	}
	received := CreateDealStore(testDeals)

	if !reflect.DeepEqual(expected, *received) {
		t.Errorf("expected %v but received %v", expected, *received)
	}
}

type dealTest struct {
	sku string
	qty  int
	expected *Deal
	expectedError bool
}

func TestGetBestDealSuccess(t *testing.T) {
	dealFor10a := Deal {"A", 10, 80}
	dealFor5a := Deal {"A", 5, 45}
	dealFor2b := Deal {"B", 2, 45}
	dealStore := DealStore {
		deals: map[string][]Deal {
			"A": {dealFor10a, dealFor5a},
			"B": {dealFor2b},
		},
	}

	tests := []dealTest {
		{"A", 7, &dealFor5a, false},
		{"A", 11, &dealFor10a, false},
		{"A", 2, nil, true},
		{"B", 2, &dealFor2b, false},
		{"B", 1, nil, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s x %d", test.sku, test.qty), func(t *testing.T) {
			received, err := dealStore.GetBestDealForSkuQty(test.sku, test.qty)

			if test.expectedError {
				if err == nil {
					t.Errorf("expected error on test %v", test)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error on test %v: %v", test, err)
				}
				if !reflect.DeepEqual(*test.expected, *received) {
					t.Errorf("expected %v but received %v", *test.expected, *received)
				}
			}
		})
	}
}
