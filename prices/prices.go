package prices

type TaxIncludedPriceJob struct {
	TaxRates          float64
	InputPrice        []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrice: []float64{10, 20, 30},
		TaxRates:   taxRate,
	}
}
