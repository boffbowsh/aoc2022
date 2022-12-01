package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main()	{
	// Open the input.txt file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxCalories := 0
	cals := 0
	// Create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the line text
		line := scanner.Text()

		if line == "" {
			// Check if the calories are higher than the max
			if cals > maxCalories {
				maxCalories = cals
			}
			cals = 0
			continue
		}

		// Parse the line as an integer
		c, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		// Add the integer to the total
		cals += c
	}

	// Print the max calories
	log.Printf("Max calories: %d", maxCalories)
}
