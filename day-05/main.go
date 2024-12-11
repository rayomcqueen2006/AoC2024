package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	rules := [][]string{}
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			rules = append(rules, []string{strings.Split(line, "|")[0], strings.Split(line, "|")[1]})
		} else {
			pass := true
			for _, rule := range rules {
				num1Index := strings.Index(line, rule[0])
				num2Index := strings.Index(line, rule[1])
				if num1Index == -1 || num2Index == -1 {
					continue
				}
				if num1Index > num2Index {
					pass = false
					break
				}
			}

			if pass {
				update := strings.Split(line, ",")
				middleIndex := int(math.Floor(float64(len(update) / 2)))
				middleNum, _ := strconv.Atoi(update[middleIndex])
				total += middleNum
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
	rules := [][]string{}
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			rules = append(rules, []string{strings.Split(line, "|")[0], strings.Split(line, "|")[1]})
		} else {
			pass := true
			for _, rule := range rules {
				num1Index := strings.Index(line, rule[0])
				num2Index := strings.Index(line, rule[1])
				if num1Index == -1 || num2Index == -1 {
					continue
				}
				if num1Index > num2Index {
					pass = false
					break
				}
			}

			if !pass {
				update := strings.Split(line, ",")
				sort.Slice(update, func(i, j int) bool {
					correctOrder := true
					for _, rule := range rules {
						if (rule[0] == update[i] || rule[0] == update[j]) && (rule[1] == update[i] || rule[1] == update[j]) {
							if rule[0] == update[j] {
								correctOrder = false
								break
							}
						}
					}
					return correctOrder
				})

				middleIndex := int(math.Floor(float64(len(update) / 2)))
				middleNum, _ := strconv.Atoi(update[middleIndex])
				total += middleNum
			}
		}
	}
	fmt.Println(total)
}
