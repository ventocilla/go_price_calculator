package main

import (
	"fmt"

	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.5}
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
