package main

import "fmt"

func threeCharWords(f string) bool {
	return len(f) == 3
}
func main() {
	animals := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	startingWithB := filter(animals, func(f string) bool {
		return f[0] == 'b'
	})
	fmt.Println(startingWithB)

	threeChar := filter(animals, threeCharWords)
	fmt.Println(threeChar)
}

func filter(animals []string, predicate func(string) bool) []string {
	var newAnimalString []string
	for _, v := range animals {
		if predicate(v) {
			newAnimalString = append(newAnimalString, v)
		}
	}
	return newAnimalString
}
