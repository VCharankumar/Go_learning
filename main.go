package main

import (
	"fmt"

	"example.com/price-calc/filemanager"
	"example.com/price-calc/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPrice(taxRate, fm)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
