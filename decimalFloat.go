package main

import (
	"fmt"
	"math"
)

func divisible(f float64) bool {
	return math.Mod(f, 3.0) == 0
}
func main() {
	floats := []float64{12.0, 9.0, 18.0, 9.65, 2.0, 27.0, 8.1, 3.3}

	integers := filter(floats, func(f float64) bool {
		return math.Mod(f, 1) == 0
	})
	fmt.Println(integers)

	divisible3 := filter(floats, divisible)
	fmt.Println(divisible3)
}

func filter(slice []float64, predicate func(float64) bool) []float64 {
	var result []float64
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}

	}
	return result
}
