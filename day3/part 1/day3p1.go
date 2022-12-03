package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// A is worth 27 points, offset by 26 (value of 'z')
const capitalCharASCIIOffset = int('A') - 26
const lowercaseCharASCIIOffset = int('a')

func main() {
	// Open file, load into input
	input, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Close input when main is finished
	defer input.Close()

	scanner := bufio.NewScanner(input)

	priorityScore := 0

	// Loop through each "token" in the scanner
	for scanner.Scan() {
		line := scanner.Text()

		lineLength := len(line)
		// We're assuming the length is always an even number
		halfPoint := lineLength / 2

		compartmentA := line[0:halfPoint]
		compartmentB := line[halfPoint:lineLength]

		// Since we know there's only one commmon character betewen the two strings
		// We can just loop through one until we find the same character in the other
		for i := 0; i < len(compartmentA); i++ {
			charAtIndexAsString := string(compartmentA[i])

			if strings.Contains(compartmentB, charAtIndexAsString) {
				priorityValue := 0
				commonCharASCIIValue := int(compartmentA[i])

				// If less than the value of 'Z', we know it's capitalized
				if commonCharASCIIValue <= int('Z') {
					priorityValue = commonCharASCIIValue - capitalCharASCIIOffset + 1
				} else {
					priorityValue = commonCharASCIIValue - lowercaseCharASCIIOffset + 1
				}

				priorityScore += priorityValue
				break
			}
		}
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%s: %d", "Priority sum is", priorityScore))
}
