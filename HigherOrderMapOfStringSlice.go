package main

import (
	"fmt"
	"strings"
)

func addingColon(word string) string {
	wordWithColon := word + ":"
	return wordWithColon
}

func main() {
	sliceOfAnimals := []string{"ant", "Beetie", "bee", "wasp", "butterfly", "fly", "moth"}

	firstWordCpitalString := Map(sliceOfAnimals, func(word string) string {
		return strings.Title(word)
	})
	fmt.Println(firstWordCpitalString)

	colonAfterWord := Map(sliceOfAnimals, addingColon)
	fmt.Println(colonAfterWord)
}

func Map(sliceOfAnimals []string, animals func(string) string) []string {
	mapString := make([]string, len(sliceOfAnimals))

	for i, val := range sliceOfAnimals {
		mapString[i] = animals(val)

	}
	return mapString
}
