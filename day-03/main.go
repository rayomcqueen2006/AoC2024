package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	r, _ := regexp.Compile("mul\\((?P<num1>[1-9][0-9]{0,2}),(?P<num2>[1-9][0-9]{0,2})\\)")

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		mem := scanner.Text()
		matches := r.FindAllStringSubmatch(mem, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			total += num1 * num2
		}
	}

	fmt.Println(total)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	r, _ := regexp.Compile("(?:do\\(\\))|(?:don't\\(\\))|(?:mul\\((?P<num1>[1-9][0-9]{0,2}),(?P<num2>[1-9][0-9]{0,2})\\))")

	scanner := bufio.NewScanner(file)
	total := 0
	do := true
	for scanner.Scan() {
		mem := scanner.Text()
		matches := r.FindAllStringSubmatch(mem, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				do = true
				continue
			}
			if match[0] == "don't()" {
				do = false
				continue
			}
			if do {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				total += num1 * num2
			}
		}
	}

	fmt.Println(total)
}
