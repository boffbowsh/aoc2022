package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main()	{
	// Open the input.txt file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elfCals := []int{}
	cals := 0
	// Create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the line text
		line := scanner.Text()

		if line == "" {
			elfCals = append(elfCals, cals)
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

	// Sort the elf calories
	sort.Ints(elfCals)
	// Add the top three together
	total := elfCals[len(elfCals)-1] + elfCals[len(elfCals)-2] + elfCals[len(elfCals)-3]
	// Print the total
	log.Println(total)
}
