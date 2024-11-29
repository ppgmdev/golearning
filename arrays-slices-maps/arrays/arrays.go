package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99} // dynamic array
	fmt.Println(prices[1])

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices) //[10.99 8.99 5.99] [10.99 8.99]

	prices = append(prices, 5.99)
	fmt.Println(prices) //[10.99 8.99 5.99]
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{10.23, 47.23, 48.10, 48.34}
	prices = append(prices, discountPrices...) //take all the elements (separated) and add them separated to the prices lists

	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 40.2, 30.1}
// 	fmt.Println(prices)

// 	productNames[2] = "A carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[3])

// 	featuredPrices := prices[1:3] //slice, take elements from element 1 to element 3 (excluded), creates new slice, does not copy original array
// 	featuredPrices2 := prices[:3] //[10.99 9.99 40.2]
// 	featuredPrices3 := prices[1:] //[9.99 40.2 30.1]

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)

// 	fmt.Println(featuredPrices)
// 	fmt.Println(featuredPrices2)
// 	fmt.Println(featuredPrices3)

// 	fmt.Println(len(prices), cap(prices))
// }
