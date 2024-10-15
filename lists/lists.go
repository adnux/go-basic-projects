package lists

import "fmt"

func StartLists() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println("productNames=>", productNames)

	prices := []float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println("prices=>", prices)
	fmt.Println("prices[2]=>", prices[2])

	featuredPrices := prices[1:]
	// featuredPrices[0] = 199.99
	featuredPrices = ReplaceSliceElement(featuredPrices, 0, 199.99)
	highlightedPrices := featuredPrices[:1]
	fmt.Println("prices=================>", prices)
	fmt.Println("featuredPrices=========>", featuredPrices)
	fmt.Println("highlightedPrices======>", highlightedPrices)
	fmt.Println("len(highlightedPrices)=>", len(highlightedPrices), "cap(highlightedPrices)=>", cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	// highlightedPrices = CutSlice(highlightedPrices, 0, 3)
	fmt.Println("highlightedPrices======>", highlightedPrices)
	fmt.Println("len(highlightedPrices)=>", len(highlightedPrices), "cap(highlightedPrices)=>", cap(highlightedPrices))

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println("prices=================>", prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println("prices=================>", prices)
}

func ReplaceSliceElement(list []float64, index int, value float64) []float64 {
	list[index] = value
	return list
}

func CutSlice(list []float64, start, end int) []float64 {
	return list[start:end]
}
