package checkout

import (
	"fmt"
	"reflect"
	"testing"
)

type checkoutTest struct {
	skus []string
	expectedTotal int
}

func getCheckoutTests() []checkoutTest {
	return []checkoutTest {
		{[]string{"A"}, 0},
	}
}

func setupCheckout() ICheckout {
	return Checkout{}
}

func TestCheckout(t *testing.T) {
	checkout := setupCheckout()
	for _, test := range getCheckoutTests() {
		t.Run(fmt.Sprintf("expected %d", test.expectedTotal), func(t *testing.T) {
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
