package prices

import (
	"fmt"

	"example.com/price_calculator/conversion"
	"example.com/price_calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRates          float64
	InputPrice        []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrice {
		taxIncludedPrice := price * (1 + job.TaxRates)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRates*100), job)

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrice: []float64{10, 20, 30},
		TaxRates:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrice = prices
}
