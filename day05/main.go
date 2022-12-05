package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Open the input.txt file and read it line by line

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Get the number of stacks
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	numStacks := (len([]rune(scanner.Text())) + 1) / 4

	// Reset
	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)

	// Get the stacks first
	stacks := make([][]rune, numStacks)
	for scanner.Scan() {
		// Read until the stack numbers
		line := scanner.Text()
		if line[1] == '1' {
			break
		}

		for i, r := range line {
			// Only look for crate letters
			if i == 0 || (i - 1) % 4 > 0 {
				continue
			}

			if r == ' ' {
				continue
			}

			// Prepend the crate to the correct stack
			stacks[(i - 1) / 4] = append([]rune{r}, stacks[(i - 1) / 4]...)
		}
	}

	// Skip a line
	scanner.Scan()

	// Get the crates to move
	for scanner.Scan() {
		// Parse the instruction, eg "move 8 from 3 to 2"
		line := scanner.Text()

		r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := r.FindStringSubmatch(line)

		// Get the number of crates to move
		numCrates, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Fatal(err)
		}

		// Get the source and destination stacks
		src, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatal(err)
		}
		dest, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatal(err)
		}

		// Move the crates
		for i := 0; i < numCrates; i++ {
			stacks[dest - 1] = append(stacks[dest - 1], stacks[src - 1][len(stacks[src - 1]) - i - 1])
		}
		stacks[src - 1] = stacks[src - 1][:len(stacks[src - 1])-numCrates]
	}

	// Print the last crate on each stack
	for _, stack := range stacks {
		fmt.Printf("%c", stack[len(stack) - 1])
	}

	fmt.Println()
}
