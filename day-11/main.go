package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitea.kandjdev.net/KandJDev/aocutils/files"
)

func main() {
	// Part One
	runStoneSim(25)
	// Part Two
	runStoneSim(75)
}

func runStoneSim(numOfBlinks int) {
	files.ScanFile("input.txt", func(line string) {
		stones := strings.Split(line, " ")
		stoneMap := map[string]int{}
		for _, stone := range stones {
			stoneMap[stone] = stoneMap[stone] + 1
		}
		for range numOfBlinks {
			stonesToExamine := make([]string, 0, len(stoneMap))
			for k := range stoneMap {
				stonesToExamine = append(stonesToExamine, k)
			}

			newStoneMap := map[string]int{}

			for _, stone := range stonesToExamine {
				if stone == "0" {
					newStoneMap["1"] += stoneMap["0"]
					continue
				}
				if len(stone)%2 == 0 {
					stone1 := strings.Join(strings.Split(stone, "")[:len(stone)/2], "")
					stone2 := strings.Join(strings.Split(stone, "")[len(stone)/2:], "")

					// convert to int and back to string to remove leading zeros
					stone1Int, _ := strconv.Atoi(stone1)
					stone1 = strconv.Itoa(stone1Int)
					stone2Int, _ := strconv.Atoi(stone2)
					stone2 = strconv.Itoa(stone2Int)

					if val, ok := newStoneMap[stone1]; ok {
						newStoneMap[stone1] = val + stoneMap[stone]
					} else {
						newStoneMap[stone1] = stoneMap[stone]
					}

					if val, ok := newStoneMap[stone2]; ok {
						newStoneMap[stone2] = val + stoneMap[stone]
					} else {
						newStoneMap[stone2] = stoneMap[stone]
					}

					continue
				}
				stoneInt, _ := strconv.Atoi(stone)
				newStoneStr := strconv.Itoa(stoneInt * 2024)
				if val, ok := newStoneMap[newStoneStr]; ok {
					newStoneMap[newStoneStr] = val + stoneMap[stone]
				} else {
					newStoneMap[newStoneStr] = stoneMap[stone]
				}
			}

			stoneMap = newStoneMap
		}
		// fmt.Println("stones: ", stoneMap)
		total := 0
		for _, num := range stoneMap {
			total += num
		}
		fmt.Println("total stones: ", total)
	})
}
