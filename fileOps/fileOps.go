package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("error reading balance")
	}
	valueText := string(data)
	parsedValue, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("error parsing Value")
	}
	return parsedValue, nil
}

func WriteFloatToFile(fileName string, value float64) {
	os.WriteFile(fileName, []byte(fmt.Sprintf("%.2f", value)), 0644)
}
