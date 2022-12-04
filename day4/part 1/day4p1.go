package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type byAssignmentPair []string

func main() {
	input, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		pair := strings.Split(line, ",")
		sort.Sort(byAssignmentPair(pair))
		firstElfAssignment := stringSliceToIntSlice(strings.Split(pair[0], "-"))
		secondElfAssignment := stringSliceToIntSlice(strings.Split(pair[1], "-"))

		if (firstElfAssignment[0]-secondElfAssignment[0]) <= 0 && (firstElfAssignment[1]-secondElfAssignment[1]) >= 0 {
			count++
		}
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%s: %d", "Assignment pairs with one range fully containing another", count))
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

func (s byAssignmentPair) Len() int {
	return len(s)
}
func (s byAssignmentPair) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byAssignmentPair) Less(i, j int) bool {
	first := stringSliceToIntSlice(strings.Split(s[i], "-"))
	second := stringSliceToIntSlice(strings.Split(s[j], "-"))

	if first[0] == second[0] {
		return first[1] > second[1]
	}
	return first[0] < second[0]
}
