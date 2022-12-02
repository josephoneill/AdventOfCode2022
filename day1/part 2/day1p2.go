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

	topThreeCaloriesCarried := [3]int{0, 0, 0}
	currElfCalorieCount := 0

	// Loop through each "token" in the scanner
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			calories, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}

			currElfCalorieCount += calories
		} else {
			// Note: It'd be far easier to add all values to a Slice, sort, and then grab the first 3 items
			// However, that wouldn't be as good for learning the nuaunces of Go

			// If the current calorie count is larger than the 3rd current largest
			if currElfCalorieCount > topThreeCaloriesCarried[2] {
				// Loop through the top three calories, find the index of where the current count falls
				for index, value := range topThreeCaloriesCarried {
					// If the current count is larger than the current value, we've found the index to insert at
					if currElfCalorieCount > value {
						// 3 - 1 as we don't need the last value
						for i := index; i < 3; i++ {
							// No need to move over last item, it will be overwritten and will result in an index of out bounds
							if i != 2 {
								topThreeCaloriesCarried[i+1] = topThreeCaloriesCarried[i]
							}

							// Apply current calorie count into it's rightful position in the top 3
							if i == index {
								topThreeCaloriesCarried[i] = currElfCalorieCount
							}
						}
						break
					}
				}
			}
			currElfCalorieCount = 0
		}
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%s: %d", "Top three highest calories count total is", findSumOfArray(topThreeCaloriesCarried[:])))
}

func findSumOfArray(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}

	return result
}
