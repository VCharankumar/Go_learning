package prices

import (
	"fmt"

	"example.com/price-calc/conversion"
	"example.com/price-calc/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	Prices            []float64           `json:"prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() error {

	err := job.loadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)

}

func (job *TaxIncludedPriceJob) loadData() error {

	lines, err := job.IOManager.ReadData()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		return err

	}

	job.Prices = prices
	return nil
}

func NewTaxIncludedPrice(taxrate float64, iomanager iomanager.IOManager) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		IOManager: iomanager,
		TaxRate:   taxrate,
		Prices:    []float64{10, 20, 30},
	}
}
