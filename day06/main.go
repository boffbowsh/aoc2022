package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	// Open the input.txt file and read it into a string
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Get the start marker
	startMarker := getStartMarker(input)
	fmt.Println(startMarker)
}

func getStartMarker(input []byte) int {
	// Loop over all characters from the 14th to the last
	for i := 14; i < len(input); i++ {
		valid := true
Outer:
		for x := i; x > i - 14; x-- {
			for y := x - 1; y > i - 14; y-- {
				if input[x] == input[y] {
					valid = false
					break Outer
				}
			}
		}
		if valid {
			return i + 1
		}
	}

	return 0
}
