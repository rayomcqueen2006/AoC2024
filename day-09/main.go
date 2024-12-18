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
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	fileSystem := []string{}
	for scanner.Scan() {
		diskMap := scanner.Text()
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

		total := 0
		for i := 0; i < len(fileSystem); i++ {
			if string(fileSystem[i]) == "." {
				continue
			}
			id, _ := strconv.Atoi(string(fileSystem[i]))
			total += (i * id)
		}

		fmt.Println(total)
	}
}
