package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AdventOfCode2022/day7/directory"
	"github.com/AdventOfCode2022/day7/file"
)

var cpuDirectory *directory.Directory
var currentDirectory = cpuDirectory
var currentPath = []string{"/"}

const directoryFileSizeThreshold = 100000.0

func main() {
	// Open file, load into input
	input, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Close input when main is finished
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		// Command
		if string(line[0]) == "$" {
			executeCommand(line)
		} else if line[0:3] == "dir" {
			_, params := getCommandAndParams(line)

			if !currentDirectory.DirectoryAlreadyExists(params) {
				addNewDirectory(params)
			}
		} else {
			// We know anything else will be a file, so add it
			fileSizeAsString, fileName := getCommandAndParams(line)
			fileSize, _ := strconv.ParseFloat(fileSizeAsString, 64)
			file := file.File{Name: fileName, Size: fileSize}

			currentDirectory.AddFile(file)
		}
	}

	// Define error handler for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := sumFileSizeGreaterThanThreshold(cpuDirectory)

	fmt.Println(fmt.Sprintf("%s: %+v", "Sum is", sum))
}

func executeCommand(line string) {
	fullCommand := line[2:]

	// If command is just ls, do nothing
	if fullCommand == "ls" {
		return
	}

	command, params := getCommandAndParams(fullCommand)

	// We know the only command other than ls is cd, perform actions for cd
	// Just to be sure (in case part 2 adds other commands), let's check
	if command == "cd" {
		if params == ".." {
			currentPath = currentPath[0 : len(currentPath)-1]
			currentDirectory = currentDirectory.ParentDirectory
			return
		} else if params == "/" {
			currentPath = currentPath[:1]

			// As long as its not our first CD into the machine (/)
			if cpuDirectory != nil {
				currentDirectory = cpuDirectory
				return
			}
		} else {
			currentPath = append(currentPath, params)
		}

		if !currentDirectory.DirectoryAlreadyExists(params) {
			currentDirectory = addNewDirectory(params)
		} else {
			currentDirectory = currentDirectory.FindDirectory(params)
		}
	}
}

func getCommandAndParams(line string) (string, string) {
	firstSpaceIndex := strings.Index(line, " ")

	command := line[:firstSpaceIndex]
	params := line[firstSpaceIndex+1:]

	return command, params
}

func addNewDirectory(name string) *directory.Directory {
	newDirectory := directory.Directory{
		Name:            name,
		FullPath:        strings.Join(currentPath[:], "/")[1:],
		Directories:     []*directory.Directory{},
		Files:           []file.File{},
		ParentDirectory: currentDirectory,
	}

	if currentDirectory != nil {
		currentDirectory.AddDirectory(&newDirectory)
	} else {
		cpuDirectory = &newDirectory
	}

	return &newDirectory
}

func sumFileSizeGreaterThanThreshold(directory *directory.Directory) float64 {
	sum := 0.0
	if len(directory.Directories) > 0 {
		for _, val := range directory.Directories {
			sum += sumFileSizeGreaterThanThreshold(val)
		}
	}

	directoryFileSize := directory.TotalDirectoryFileSize()
	if directoryFileSize <= directoryFileSizeThreshold {
		sum += directoryFileSize
	}

	return sum
}
