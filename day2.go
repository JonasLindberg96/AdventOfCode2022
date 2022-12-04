package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("lib/rps.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	totalPoints := 0
	secondPartPoints := 0

	for fileScanner.Scan() {
		opponent := string(fileScanner.Text()[0])
		response := string(fileScanner.Text()[2])

		// First part

		// Add Shape Points
		if response == "X" { // Rock
			totalPoints = totalPoints + 1
		} else if response == "Y" { // Paper
			totalPoints = totalPoints + 2
		} else { // Scissors
			totalPoints = totalPoints + 3
		}

		// Add Outcome Points
		if opponent == "A" { // Opponent plays Rock
			if response == "X" { // Response is Rock
				totalPoints = totalPoints + 3 // Draw
			} else if response == "Y" { // Response is Paper
				totalPoints = totalPoints + 6 // Win
			}
		} else if opponent == "B" { // Opponent plays Paper
			if response == "Y" { // Response is Paper
				totalPoints = totalPoints + 3 // Draw
			} else if response == "Z" { // Response is Scissors
				totalPoints = totalPoints + 6 // Win
			}
		} else { // Opponent plays Scissors
			if response == "Z" { // Response is Scissors
				totalPoints = totalPoints + 3
			} else if response == "X" { // Response is Rock
				totalPoints = totalPoints + 6
			}
		}

		// Second Part
		if opponent == "A" { // Opponent plays Rock
			if response == "X" { // Need to Lose (Play Scissors)
				secondPartPoints = secondPartPoints + 0 + 3
			} else if response == "Y" { // Need to Draw (Play Rock)
				secondPartPoints = secondPartPoints + 3 + 1
			} else { // Need to Win (Play Paper)
				secondPartPoints = secondPartPoints + 6 + 2
			}
		} else if opponent == "B" { // Opponents plays Paper
			if response == "X" { // Need to Lose (Play Rock)
				secondPartPoints = secondPartPoints + 0 + 1
			} else if response == "Y" { // Need to Draw (Play Paper)
				secondPartPoints = secondPartPoints + 3 + 2
			} else { // Need to Win (Play Scissors)
				secondPartPoints = secondPartPoints + 6 + 3
			}
		} else { // Opponent plays Scissors
			if response == "X" { // Need to Lose (Play Paper)
				secondPartPoints = secondPartPoints + 0 + 2
			} else if response == "Y" { // Need to Draw (Play Scissors)
				secondPartPoints = secondPartPoints + 3 + 3
			} else { // Need to Win (Play Rock)
				secondPartPoints = secondPartPoints + 6 + 1
			}
		}
	}

	readFile.Close()

	fmt.Println("First part points: ", totalPoints)
	fmt.Println("Second part point: ", secondPartPoints)
}

/*
A = Rock
B = Paper
C = Scissors

X = Rock
Y = Paper
Z = Scissors

Rock defeats Scissors
Scissors defeats Paper
Paper defeats Rock

Shape Points:
Rock = 1
Paper = 2
Scissors = 3

Outcome Points:
Loss = 0
Draw = 3
Win = 6
*/
