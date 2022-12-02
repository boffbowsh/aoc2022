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
	var strategy []string
  for scanner.Scan() {
    strategy = append(strategy, scanner.Text())
  }

  // Simulate each round of the game
  score := 0
  for _, round := range strategy {
    // Parse the round to get the opponent's move and your recommended move
    parts := strings.Split(round, " ")
    opponentMove := parts[0]
    recommendedMove := parts[1]

    // Simulate the round
    outcome := simulateRound(opponentMove, recommendedMove)

    // Calculate the score for the round
    roundScore := getScore(recommendedMove, outcome)

    // Add the round score to the total score
    score += roundScore
  }

  // Print the total score
  fmt.Println(score)
}

func simulateRound(opponentMove, recommendedMove string) string {
  // Define the rules for the game
  rules := map[string]map[string]string{
    "A": map[string]string{
      "X": "draw",
      "Y": "lose",
      "Z": "win",
    },
    "B": map[string]string{
      "X": "win",
      "Y": "draw",
      "Z": "lose",
    },
    "C": map[string]string{
      "X": "lose",
      "Y": "win",
      "Z": "draw",
    },
  }

  // Return the outcome of the round
  return rules[opponentMove][recommendedMove]
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
  case "win":
    score += 6
  case "draw":
    score += 3
  }

  return score
}
