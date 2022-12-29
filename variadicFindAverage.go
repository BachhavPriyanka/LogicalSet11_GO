package main

import "fmt"

func main() {
	computeAvg := average(45, 67, 80, 71, 60)
	fmt.Println(computeAvg)
}

func average(numbers ...int) float64 {
	var add int
	var count int = 0
	for _, val := range numbers {
		add += val
		count++
	}
	avg := float64(add) / float64(count)
	return avg
}
