package scanner

import (
	"checkout/pricing"
	"fmt"
)

// Interface to adhere to
type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct {
	prices map[string]int
	deals *pricing.DealStore
	basket []string
}

func CreateCheckout(prices map[string]int, deals *pricing.DealStore) Checkout {
	return Checkout{prices, deals, []string{}}
}

func (c *Checkout) Scan(item string) {
	if _, ok := c.prices[item]; ok {
		c.basket = append(c.basket, item)
		fmt.Println("OK")
	} else {
		fmt.Printf("Unknown item %s\n", item)
	}
}

func (c *Checkout) GetTotalPrice() int {
	total := 0
	basketBySku := getBasketBySku(c.basket)
	for sku, qty := range basketBySku {
		total += c.calculateTotalForSku(sku, qty)
	}
	return total
}

func (c Checkout) calculateTotalForSku(sku string, qty int) int {
	total := 0
	leftToProcess := qty
	for leftToProcess > 0 {
		deal, err := c.deals.GetBestDealForSkuQty(sku, leftToProcess)
		if err != nil {
			total += leftToProcess * c.prices[sku]
			leftToProcess = 0
		} else {
			total += deal.Price
			leftToProcess -= deal.Qty
		}
	}
	return total
}

func getBasketBySku(basket []string) map[string]int {
	basketBySku := map[string]int{}
	for _, sku := range basket {
		totalForSku, ok := basketBySku[sku]
		if ok {
			basketBySku[sku] = totalForSku + 1
		} else {
			basketBySku[sku] = 1
		}
	}
	return basketBySku
}
