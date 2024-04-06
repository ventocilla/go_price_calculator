package prices

type TaxIncludedPriceJob struct {
	TaxRates          float64
	InputPrice        []float64
	TaxIncludedPrices map[string]float64
}
