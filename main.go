package main

import (
	"example.com/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.5}
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
