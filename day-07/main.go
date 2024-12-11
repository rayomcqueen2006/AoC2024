package main

import (
	"bufio"
	"fmt"
	"math"
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
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		targetNum, _ := strconv.Atoi(strings.Split(line, ":")[0])
		operandsStr := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
		operands := []int{}
		for _, operand := range operandsStr {
			num, _ := strconv.Atoi(operand)
			operands = append(operands, num)
		}

		permutations := calculatePermuatations(len(operands)-1, 2, []string{"+", "*"})
		for _, permutation := range permutations {
			currentResult := operands[0]
			for i := 1; i < len(operands); i++ {
				if string(permutation[i-1]) == "+" {
					currentResult = currentResult + operands[i]
				} else if string(permutation[i-1]) == "*" {
					currentResult = currentResult * operands[i]
				}
			}
			if currentResult == targetNum {
				total += targetNum
				break
			}
		}
	}

	fmt.Println(total)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		targetNum, _ := strconv.Atoi(strings.Split(line, ":")[0])
		operandsStr := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
		operands := []int{}
		for _, operand := range operandsStr {
			num, _ := strconv.Atoi(operand)
			operands = append(operands, num)
		}

		permutations := calculatePermuatations(len(operands)-1, 3, []string{"+", "*", "|"})
		for _, permutation := range permutations {
			currentResult := operands[0]
			for i := 1; i < len(operands); i++ {
				if string(permutation[i-1]) == "+" {
					currentResult = currentResult + operands[i]
				} else if string(permutation[i-1]) == "*" {
					currentResult = currentResult * operands[i]
				} else if string(permutation[i-1]) == "|" {
					currentResult, _ = strconv.Atoi(strconv.Itoa(currentResult) + strconv.Itoa(operands[i]))
				}
			}
			if currentResult == targetNum {
				total += targetNum
				break
			}
		}
	}

	fmt.Println(total)
}

func calculatePermuatations(size int, base int, chars []string) []string {
	maxNum := math.Pow(float64(base), float64(size))
	fmtString := "%0" + strconv.Itoa(size) + "s"
	permutations := []string{}
	for i := 0; i < int(maxNum); i++ {
		permutation := fmt.Sprintf(fmtString, strconv.FormatInt(int64(i), base))
		for j := 0; j < base; j++ {
			permutation = strings.ReplaceAll(permutation, strconv.Itoa(j), chars[j])
		}
		permutations = append(permutations, permutation)
	}
	return permutations
}
