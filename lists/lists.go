package lists

import "fmt"

func StartLists() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println("productNames=>", productNames)

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println("prices=>", prices)
	fmt.Println("prices[2]=>", prices[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println("prices=================>", prices)
	fmt.Println("featuredPrices=========>", featuredPrices)
	fmt.Println("highlightedPrices======>", highlightedPrices)
	fmt.Println("len(highlightedPrices)=>", len(highlightedPrices), "cap(highlightedPrices)=>", cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println("highlightedPrices=>", highlightedPrices)
	fmt.Println("len(highlightedPrices)=>", len(highlightedPrices), "cap(highlightedPrices)=>", cap(highlightedPrices))
}
