package scanner

import (
	"fmt"
	"github.com/Thustle/checkout/pricing"
	"reflect"
	"testing"
)

type checkoutTest struct {
	skus          []string
	expectedTotal int
}

func getCheckoutTests() []checkoutTest {
	return []checkoutTest{
		{[]string{"A"}, 50},
		{[]string{"A", "A", "A"}, 130},
		{[]string{"A", "A", "A", "A"}, 180},
		{[]string{"A", "A", "A", "A", "A"}, 230},
		{[]string{"A", "A", "A", "A", "A", "A"}, 260},
		{[]string{}, 0},
		{[]string{"B", "B"}, 45},
		{[]string{"A", "A", "A", "B", "B"}, 175},
	}
}

func setupCheckout() Checkout {
	prices := map[string]int{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}
	deals := map[string][]pricing.Deal{
		"A": {{"A", 3, 130}},
		"B": {{"B", 2, 45}},
	}
	dataStore := pricing.CreateDealStore(deals)
	return CreateCheckout(prices, dataStore)
}

func TestCheckout(t *testing.T) {
	for _, test := range getCheckoutTests() {
		t.Run(fmt.Sprintf("expected %d", test.expectedTotal), func(t *testing.T) {
			checkout := setupCheckout()
			for _, scan := range test.skus {
				checkout.Scan(scan)
			}
			received := checkout.GetTotalPrice()
			if !reflect.DeepEqual(test.expectedTotal, received) {
				t.Errorf("expected %v but received %v", test.expectedTotal, received)
			}
		})
	}
}

func TestCreateCheckout(t *testing.T) {
	prices := map[string]int{"A": 1}
	dealStore := pricing.DealStore{}

	expected := Checkout{prices, &dealStore, []string{}}
	received := CreateCheckout(prices, &dealStore)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("expected %v but received %v", expected, received)
	}
}
