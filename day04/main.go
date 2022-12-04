package main

import (
	"fmt"
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

        start1, end1 := assignment1[0], assignment1[1]
        start2, end2 := assignment2[0], assignment2[1]

        if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
            overlappingAssignments++
        }
    }

    return overlappingAssignments
}

func main() {
    input := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
    fmt.Println(findOverlappingAssignments(input)) // expected output: 2
}
