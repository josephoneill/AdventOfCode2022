package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const oppositionCharASCIIOffset = int('A')

func main() {
	// Open file, load into input
	input, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Close input when main is finished
	defer input.Close()

	scanner := bufio.NewScanner(input)

	score := 0

	// Loop through each "token" in the scanner
	for scanner.Scan() {
		line := scanner.Text()

		//  Splits the string. rpsRound[0] is other oppositionMove's move, rpsRound[1] is my move
		rpsRound := strings.Split(line, " ")

		oppositionMove := int(rpsRound[0][0]) - oppositionCharASCIIOffset
		desiredOutcome := rpsRound[1]
		myMove := 0

		switch {
		case desiredOutcome == "X":
			myMove = int(mod(oppositionMove-1, 3))
		case desiredOutcome == "Y":
			myMove = oppositionMove
			score += 3
		case desiredOutcome == "Z":
			myMove = int(mod(oppositionMove+1, 3))
			score += 6
		}

		score += myMove + 1
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%s: %d", "Total score is", score))
}

func mod(a, b int) int {
	return (a%b + b) % b
}
