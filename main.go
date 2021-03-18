package main

import (
	"bufio"
	"checkout/pricing"
	"checkout/scanner"
	"fmt"
	"log"
	"os"
)

const (
	dealsFile  = "deals-sku_qty_price.csv"
	pricesFile = "prices-sku_price.csv"
)

func main() {
	checkout := getCheckoutScanner()

	itemScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Scan Item > ")
		itemScanner.Scan()
		text := itemScanner.Text()
		if len(text) != 0 {
			checkout.Scan(text)
		} else {
			fmt.Printf("Total: %d\n", checkout.GetTotalPrice())
			break
		}
	}
}

func getCheckoutScanner() scanner.Checkout {
	file, err := os.Open(dealsFile)
	if err != nil {
		log.Fatalln("Could not open deals file", err)
	}
	dealStore, err := pricing.LoadDeals(file)
	if err != nil {
		log.Fatalln("Unable to get deals", err)
	}

	file, err = os.Open(pricesFile)
	if err != nil {
		log.Fatalln("Could not open deals file", err)
	}
	prices, err := pricing.LoadPrices(file)
	if err != nil {
		log.Fatalln("Unable to get prices", err)
	}
	return scanner.CreateCheckout(*prices, dealStore)
}
