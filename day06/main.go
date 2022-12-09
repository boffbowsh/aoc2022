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

	// Loop over all characters from the 4th to the last
	for i := 3; i < len(input); i++ {
		// Check if the last 4 characters are all different
		if input[i] != input[i-1] &&
			input[i] != input[i-2] &&
			input[i] != input[i-3] &&
			input[i-1] != input[i-2] &&
			input[i-1] != input[i-3] &&
			input[i-2] != input[i-3] {
			// Print the offset
			fmt.Printf("%d: %c\n", i+1, input[i-3:i+1])
			break
		}
	}
}
