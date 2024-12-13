package main

import (
	"fmt"

	"example.com/calculator/filemanager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])

		//if err != nil {
		//fmt.Println("could not process job")
		//fmt.Println(err)
		//}

	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
