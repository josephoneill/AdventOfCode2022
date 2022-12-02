package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file, load into input
	input, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Close input when main is finished
	defer input.Close()

	scanner := bufio.NewScanner(input)

	currElfCalorieCount := 0
	highestCalorieCount := 0

	// Loop through each "token" in the scanner
	for scanner.Scan() {
		line := scanner.Text()

		// If the line isn't blank, we're on the same elf
		if line != "" {
			// Convert string to number
			calories, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}

			currElfCalorieCount += calories
		} else {
			// We've finished the current elf, check if their count is higher than current highest
			if currElfCalorieCount > highestCalorieCount {
				highestCalorieCount = currElfCalorieCount
			}
			currElfCalorieCount = 0
		}
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%s: %d", "Highest calorie count is", highestCalorieCount))
}
