# Checkout

## Description
A simple command line app to calculate checkout totals.
The app loads external pricing and deal information and 
uses these to determine the final total.

The external pricing and deal files are:
- prices-sku_price.csv (which contains sku and price fields)
- deals-sku_qty_price.csv (which contains sku, qty and price fields)

## Usage
Once running, the program will prompt you for an item SKU. 
Enter one and press return. You should see an OK message if the 
SKU has been recognised and you will be able to enter the next SKU.
Once all items have been scanned just press the return key to
see the total.