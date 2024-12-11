package main

import (
	"bufio"
	"fmt"
	"os"
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

	words := []string{"XMAS", "SAMX"}

	wordSearch := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		wordSearch = append(wordSearch, strings.Split(line, ""))
	}

	wordCount := 0

	// Horizontal
	for _, hLine := range wordSearch {
		wordCount += getWordCount(words, hLine)
	}

	// Vertical
	for i := 0; i < len(wordSearch[0]); i++ {
		vLine := []string{}
		for j := 0; j < len(wordSearch); j++ {
			vLine = append(vLine, wordSearch[j][i])
		}
		wordCount += getWordCount(words, vLine)
	}

	// Diagonal 1
	for i := 0; i < len(wordSearch[0]); i++ {
		dLine := []string{}
		step := 0
		for i+step < len(wordSearch[0]) && step < len(wordSearch) {
			dLine = append(dLine, wordSearch[0+step][i+step])
			step += 1
		}
		wordCount += getWordCount(words, dLine)
	}

	for i := 1; i < len(wordSearch); i++ {
		dLine := []string{}
		step := 0
		for i+step < len(wordSearch) && step < len(wordSearch[0]) {
			dLine = append(dLine, wordSearch[i+step][0+step])
			step += 1
		}
		wordCount += getWordCount(words, dLine)
	}

	// Diagonal 2
	for i := 0; i < len(wordSearch[0]); i++ {
		dLine := []string{}
		step := 0
		for i-step >= 0 && step < len(wordSearch) {
			dLine = append(dLine, wordSearch[0+step][i-step])
			step += 1
		}
		wordCount += getWordCount(words, dLine)
	}

	for i := 1; i < len(wordSearch); i++ {
		dLine := []string{}
		step := 0
		for len(wordSearch[0])-step >= 0 && i+step < len(wordSearch) {
			dLine = append(dLine, wordSearch[i+step][len(wordSearch[0])-step-1])
			step += 1
		}
		wordCount += getWordCount(words, dLine)
	}

	fmt.Println(wordCount)
}

func getWordCount(words []string, line []string) int {
	lineStr := strings.Join(line, "")
	totalCount := 0
	for _, word := range words {
		j := 0
		count := 0
		for {
			i := strings.Index(lineStr[j:], word)
			if i == -1 {
				break
			}
			count += 1
			j += i + 1
		}
		totalCount += count
	}
	return totalCount
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	wordSearch := [][]string{}
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		wordSearch = append(wordSearch, strings.Split(line, ""))
	}

	for y := 0; y < len(wordSearch); y++ {
		for x := 0; x < len(wordSearch[0]); x++ {
			if x+2 >= len(wordSearch[0]) || y+2 >= len(wordSearch) {
				continue
			}
			word1 := wordSearch[y][x] + wordSearch[y+1][x+1] + wordSearch[y+2][x+2]
			word2 := wordSearch[y][x+2] + wordSearch[y+1][x+1] + wordSearch[y+2][x]
			if (word1 == "MAS" || word1 == "SAM") && (word2 == "MAS" || word2 == "SAM") {
				count += 1
			}
		}
	}

	fmt.Println(count)
}
