package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  // Read the strategy guide from input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

  score := 0

  for scanner.Scan() {
    round := scanner.Text()

    // Parse the round to get the opponent's move and your recommended move
    parts := strings.Split(round, " ")
    opponentMove := parts[0]
    desiredOutcome := parts[1]

    // Simulate the round
    recommendedMove := simulateRound(opponentMove, desiredOutcome)

    // Calculate the score for the round
    roundScore := getScore(recommendedMove, desiredOutcome)

    // Add the round score to the total score
    score += roundScore
  }

  // Print the total score
  fmt.Println(score)
}

func simulateRound(opponentMove, desiredOutcome string) string {
  // Define the rules for the game
  rules := map[string]map[string]string{
    "A": {
      "Y": "A",
      "Z": "B",
      "X": "C",
    },
    "B": {
      "X": "A",
      "Y": "B",
      "Z": "C",
    },
    "C": {
      "Z": "A",
      "X": "B",
      "Y": "C",
    },
  }

  // Return the outcome of the round
  return rules[opponentMove][desiredOutcome]
}

func getScore(move, outcome string) int {
  // Define the scores for each move
  scores := map[string]int{
    "A": 1,
		"B": 2,
		"C": 3,
  }

  // Calculate the score for the round
  score := scores[move]
  switch outcome {
  case "Z":
    score += 6
  case "Y":
    score += 3
  }

  return score
}
