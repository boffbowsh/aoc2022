package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// findOverlappingAssignments takes in a slice of strings representing the
// section assignment pairs, and returns the number of pairs that have one
// assignment fully contained by the other.
func findOverlappingAssignments(input []string) int {
    overlappingAssignments := 0

    for _, pair := range input {
        assignments := strings.Split(pair, ",")
        assignment1 := strings.Split(assignments[0], "-")
        assignment2 := strings.Split(assignments[1], "-")

        start1, _ := strconv.Atoi(assignment1[0])
        end1, _ := strconv.Atoi(assignment1[1])
        start2, _ := strconv.Atoi(assignment2[0])
        end2, _ := strconv.Atoi(assignment2[1])

        if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
						overlappingAssignments++
				}
    }

    return overlappingAssignments
}

func main() {
		// Open the input.txt file.
		file, err := os.Open("input.txt")
		if err != nil {
				log.Fatal(err)
		}
		defer file.Close()

		// Read each line into a slice of strings.
		var input []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
				input = append(input, scanner.Text())
		}

    fmt.Println(findOverlappingAssignments(input)) // expected output: 2
}
