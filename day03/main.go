package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Open the input.txt file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to store the lines
	var rucksacks []string

	// Loop over the lines in the file
	for scanner.Scan() {
		// Append the line to the slice
		rucksacks = append(rucksacks, scanner.Text())
	}

  // Map to count the number of times each character appears in both compartments of each rucksack
  charCount := make(map[rune]int)

  for _, rucksack := range rucksacks {
    // Split the rucksack into its two compartments
    first := rucksack[:len(rucksack)/2]
    second := rucksack[len(rucksack)/2:]

    // Go through each character in the first compartment and check if it appears in the second compartment
    for _, char := range first {
      if strings.ContainsRune(second, char) {
        charCount[char]++
      }
    }
  }

  // The sum of the priorities of the item types that appear in both compartments of each rucksack
  var sum int

  // Go through each character and its count in the map
  for char, count := range charCount {
    // If the character appears in both compartments of every rucksack, add its priority to the sum
    if count == len(rucksacks) {
      // Lowercase characters have priorities 1 through 26
      if char >= 'a' && char <= 'z' {
        sum += int(char) - 'a' + 1
      } else {
        // Uppercase characters have priorities 27 through 52
        sum += int(char) - 'A' + 27
      }
    }
  }

  fmt.Println(sum)
}
