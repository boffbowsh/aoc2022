package main

import "testing"

func TestFindOverlappingAssignments(t *testing.T) {
    // Test the example in the prompt.
    input := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"}
    expected := 2
    actual := findOverlappingAssignments(input)
    if expected != actual {
        t.Errorf("Test case 1 failed: expected %d, got %d", expected, actual)
    }

    // Test a case with no overlapping assignments.
    input = []string{"1-10,11-20", "21-30,31-40", "41-50,51-60", "61-70,71-80", "81-90,91-100"}
    expected = 0
    actual = findOverlappingAssignments(input)
    if expected != actual {
        t.Errorf("Test case 2 failed: expected %d, got %d", expected, actual)
    }

    // Test a case with overlapping assignments between every pair of Elves.
    input = []string{"1-50,1-50", "1-50,1-50", "1-50,1-50", "1-50,1-50", "1-50,1-50", "1-50,1-50"}
    expected = 15
    actual = findOverlappingAssignments(input)
    if expected != actual {
        t.Errorf("Test case 3 failed: expected %d, got %d", expected, actual)
    }

    // Test a case with all assignments overlapping with each other.
    input = []string{"1-100,1-100", "1-100,1-100", "1-100,1-100", "1-100,1-100", "1-100,1-100", "1-100,1-100"}
    expected = 15
    actual = findOverlappingAssignments(input)
    if expected != actual {
        t.Errorf("Test case 4 failed: expected %d, got %d", expected, actual)
    }
}
