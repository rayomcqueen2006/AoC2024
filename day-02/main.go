package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mode int

const (
	INCREASING Mode = iota
	DECREASING
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

	var mode Mode
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")

		valid := true

		for i, level := range levels {
			if i == len(levels)-1 {
				break
			}

			currentLevel, _ := strconv.Atoi(level)
			nextLevel, _ := strconv.Atoi(levels[i+1])

			if i == 0 {
				if nextLevel-currentLevel > 0 {
					mode = INCREASING
				} else if currentLevel-nextLevel > 0 {
					mode = DECREASING
				} else {
					valid = false
					break
				}
			}

			if mode == INCREASING {
				if !(nextLevel-currentLevel >= 1 && nextLevel-currentLevel <= 3) {
					valid = false
					break
				}
			} else {
				if !(currentLevel-nextLevel >= 1 && currentLevel-nextLevel <= 3) {
					valid = false
					break
				}
			}
		}

		if valid {
			total += 1
		}
	}

	fmt.Println(total)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")

		valid := false

		for i := 0; i <= len(levels)+1; i++ {
			ogLevels := strings.Split(report, " ")
			validIteration := checkValid(levels)

			if validIteration {
				valid = true
				break
			}

			if i > len(levels) {
				break
			}
			levels = append(ogLevels[:i], ogLevels[i+1:]...)
		}

		if valid {
			total += 1
		}
	}

	fmt.Println(total)
}

func checkValid(levels []string) bool {
	var mode Mode
	valid := true

	for i, level := range levels {
		if i == len(levels)-1 {
			break
		}

		currentLevel, _ := strconv.Atoi(level)
		nextLevel, _ := strconv.Atoi(levels[i+1])

		if i == 0 {
			if nextLevel-currentLevel > 0 {
				mode = INCREASING
			} else if currentLevel-nextLevel > 0 {
				mode = DECREASING
			} else {
				valid = false
				break
			}
		}

		if mode == INCREASING {
			if !(nextLevel-currentLevel >= 1 && nextLevel-currentLevel <= 3) {
				valid = false
				break
			}
		} else {
			if !(currentLevel-nextLevel >= 1 && currentLevel-nextLevel <= 3) {
				valid = false
				break
			}
		}
	}

	return valid
}
