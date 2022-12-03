package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

  var sum int

  // Loop over the lines in the file
  for scanner.Scan() {
    set1 := lineToSet(scanner.Text())
    scanner.Scan()
    set2 := lineToSet(scanner.Text())
    scanner.Scan()
    set3 := lineToSet(scanner.Text())
    intersectionSet := setIntersection(set1, set2, set3)

    // Get only element in the set. Error if there's more than one element
    var char rune
    for char = range intersectionSet {
      break
    }

    sum += charPriority(char)
  }

  fmt.Println(sum)
}

func charPriority(char rune) int {
  if char >= 'a' && char <= 'z' {
    return int(char) - 'a' + 1
  } else {
    return int(char) - 'A' + 27
  }
}

func lineToSet(line string) map[rune]bool {
  set := make(map[rune]bool)
  for _, char := range line {
    set[char] = true
  }
  return set
}

func setIntersection(sets ...map[rune]bool) map[rune]bool {
  intersectionSet := make(map[rune]bool)
  for char := range sets[0] {
    intersectionSet[char] = true
  }

  for _, set := range sets {
    for char := range intersectionSet {
      if !set[char] {
        delete(intersectionSet, char)
      }
    }
  }
  return intersectionSet
}
