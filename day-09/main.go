package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diskMap := scanner.Text()
		fileSystem := generateFilesystem(diskMap)

		endPointer := len(fileSystem) - 1

		for i := 0; i < len(fileSystem); i++ {
			if fileSystem[i] == "." {
				// find last number
				for string(fileSystem[endPointer]) == "." {
					endPointer -= 1
				}
				if endPointer <= i {
					break
				}
				fileSystem[i] = fileSystem[endPointer]
				fileSystem[endPointer] = "."
			}
		}

		fmt.Println(calculateChecksum(fileSystem))
	}
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	movedOrAttemptedIds := make(map[string]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		diskMap := scanner.Text()
		fileSystem := generateFilesystem(diskMap)

		endPointer := len(fileSystem) - 1

		for endPointer >= 0 {
			// Find next file block
			if fileSystem[endPointer] == "." {
				endPointer--
				continue
			}
			if endPointer < 0 {
				break
			}

			fileId := fileSystem[endPointer]

			// Find length of block
			fileLen := 0
			tmpPointer := endPointer
			for tmpPointer >= 0 && fileSystem[tmpPointer] == fileId {
				fileLen++
				tmpPointer--
			}

			// If we've attempted to move this file before, continue
			if _, ok := movedOrAttemptedIds[fileId]; ok {
				endPointer = endPointer - fileLen
				continue
			}

			// Record that we're attempting to move the file
			movedOrAttemptedIds[fileId] = struct{}{}

			// Find the first set of free space
			i := 0
			match := false
			freeSpace := 0
			for i <= endPointer && !match {
				if fileSystem[i] == "." {
					freeSpace++
					if freeSpace == fileLen {
						match = true
					}
				} else {
					freeSpace = 0
				}
				i++
			}

			if match {
				// We can move the file into the free space
				startingIndex := i - fileLen
				for idx := range fileLen {
					fileSystem[startingIndex+idx] = fileId
					fileSystem[endPointer-idx] = "."
				}
			}

			endPointer = endPointer - fileLen
		}

		fmt.Println(calculateChecksum(fileSystem))
	}
}

func generateFilesystem(diskMap string) []string {
	fileSystem := []string{}
	id := 0
	for i, length := range strings.Split(diskMap, "") {
		maxLen, _ := strconv.Atoi(length)
		if i%2 == 0 {
			// this is a file
			for j := 0; j < maxLen; j++ {
				fileSystem = append(fileSystem, strconv.Itoa(id))
			}
			id += 1
		} else {
			// this is empty space
			for j := 0; j < maxLen; j++ {
				fileSystem = append(fileSystem, ".")
			}
		}
	}
	return fileSystem
}

func calculateChecksum(fileSystem []string) int {
	total := 0
	for i := 0; i < len(fileSystem); i++ {
		if string(fileSystem[i]) == "." {
			continue
		}
		id, _ := strconv.Atoi(string(fileSystem[i]))
		total += (i * id)
	}
	return total
}
