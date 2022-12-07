package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const slidingWindowPaneSize = 14

func main() {
	// Open file, load into input
	input, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Close input when main is finished
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	line := scanner.Text()

	bufferCount := slidingWindowPaneSize
	fullWindowSize := len(line)
	slidingWindowContent := ""
	slidingWindowContent = line[0:slidingWindowPaneSize]
	foundStartOfMarker := containsUniqueCharacters(slidingWindowContent)

	for i := slidingWindowPaneSize; i < fullWindowSize && !foundStartOfMarker; i++ {
		bufferCount++
		slidingWindowContent = slidingWindowContent[1:]
		slidingWindowContent += string(line[i])
		foundStartOfMarker = containsUniqueCharacters(slidingWindowContent)
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%s: %d", "Characters processed before start-of-packet marker", bufferCount))
}

func containsUniqueCharacters(s string) bool {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) > 1 {
			return false
		}
	}

	return true
}
