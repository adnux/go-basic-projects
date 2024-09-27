package maps

import (
	"fmt"

	"github.com/adnux/go-basic-projects/maps/maps"
)

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println("courseRatings=>", m)
}

func MapsArrays() {
	// this works
	userNames := make([]string, 2, 5)
	// this doesn't work
	// userNames := []string{}

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println("userNames=>", userNames)

	// courseRatings := map[string]float64{}
	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("Index:", index, "=> Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key, "=> Value:", value)
	}

	maps.MapsWebsite()
}
