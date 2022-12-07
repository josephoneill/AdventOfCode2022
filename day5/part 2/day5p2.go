package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const crateStrWidth = 4.0

func main() {
	// Open file, load into input
	input, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Close input when main is finished
	defer input.Close()

	scanner := bufio.NewScanner(input)

	filterOutNumbersRegex := regexp.MustCompile("[0-9]+")
	stackingInitialCrates := true
	var crateStacks [][]string
	var numOfColumns int

	// Loop through each "token" in the scanner
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			stackingInitialCrates = false
			continue
		}

		if stackingInitialCrates {
			if len(crateStacks) == 0 {
				numOfColumns = int(math.Ceil(float64(len(line)) / crateStrWidth))

				crateStacks = make([][]string, numOfColumns)
				for i := range crateStacks {
					crateStacks[i] = []string{}
				}
			}

			// If the line does not contain [, we've hit the column numbering line, skip
			if !strings.Contains(line, "[") {
				continue
			}

			for i := 0; i < numOfColumns; i++ {
				substrIndex := i * crateStrWidth
				startingIndex := substrIndex

				crate := strings.TrimSpace(line[startingIndex : startingIndex+crateStrWidth-1])

				if crate == "" {
					continue
				}

				// We know due to the strict format of the input that the content of
				// the crate will be one single character at index 1 of our "crate" string
				crateStacks[i] = append(crateStacks[i], string(crate[1]))
			}
		} else {
			// []string contains move count, move source column, and move destination column, respectively
			crateMoveInstructions := stringSliceToIntSlice(filterOutNumbersRegex.FindAllString(line, -1))
			// Subtract 1 because indexes start at 0 but columns start at 1
			for i := 1; i < 3; i++ {
				crateMoveInstructions[i] = crateMoveInstructions[i] - 1
			}

			moveCrateCount := crateMoveInstructions[0]
			moveSrcColumn := crateMoveInstructions[1]
			moveSrcDest := crateMoveInstructions[2]

			for i := moveCrateCount - 1; i >= 0; i-- {
				crateStacks[moveSrcDest] = prependString(crateStacks[moveSrcDest], crateStacks[moveSrcColumn][i])
				crateStacks[moveSrcColumn] = removeIndex(crateStacks[moveSrcColumn], i)
			}
		}
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sb strings.Builder
	for _, crateCol := range crateStacks {
		sb.WriteString(crateCol[0])
	}

	fmt.Println(fmt.Sprintf("%s: %s", "Crates appear as", sb.String()))
}

func stringSliceToIntSlice(s []string) []int {
	var intSlice = []int{}

	for _, x := range s {
		val, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, val)
	}

	return intSlice
}

func removeIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// https://stackoverflow.com/questions/53737435/how-to-prepend-int-to-slice
func prependString(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y
	return x
}
