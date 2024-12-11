package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	antennas := map[string][][]int{}
	grid := [][]string{}
	total := 0

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for x, symbol := range line {
			if string(symbol) != "." {
				existingCoords, pres := antennas[string(symbol)]
				if !pres {
					antennas[string(symbol)] = [][]int{{x, y}}
				} else {
					existingCoords = append(existingCoords, []int{x, y})
					antennas[string(symbol)] = existingCoords
				}
			}
			row = append(row, string(symbol))
		}
		grid = append(grid, row)
		y += 1
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			// get distances from each antenna type
			for _, coords := range antennas {
				distances := [][]int{}
				antennaMatch := false
				for _, coord := range coords {
					distances = append(distances, []int{coord[0] - x, coord[1] - y})
				}
				// check if any distances are multiples of 2 of any other distances
				for index, distance := range distances {
					match := false
					if distance[0] == 0 && distance[1] == 0 {
						continue
					}
					for i := 0; i < len(distances); i++ {
						if i == index {
							continue
						}
						if (distance[0]*2 == distances[i][0]) && (distance[1]*2 == distances[i][1]) {
							match = true
							break
						}
					}
					if match {
						antennaMatch = true
						total += 1
						break
					}
				}
				if antennaMatch {
					break
				}
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

	antennas := map[string][][]int{}
	grid := [][]string{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for x, symbol := range line {
			if string(symbol) != "." {
				existingCoords, pres := antennas[string(symbol)]
				if !pres {
					antennas[string(symbol)] = [][]int{{x, y}}
				} else {
					existingCoords = append(existingCoords, []int{x, y})
					antennas[string(symbol)] = existingCoords
				}
			}
			row = append(row, string(symbol))
		}
		grid = append(grid, row)
		y += 1
	}

	antinodes := map[string]struct{}{}

	for _, coords := range antennas {
		for index, coord := range coords {
			for i := 0; i < len(coords); i++ {
				if index == i {
					continue
				}
				// calculate line equation
				m := float64(coord[1]-coords[i][1]) / float64(coord[0]-coords[i][0])
				c := float64(coord[1]) - (m * float64(coord[0]))

				// check if any points match the line equation
				for y := 0; y < len(grid); y++ {
					for x := 0; x < len(grid[0]); x++ {
						value := (m * float64(x)) + c
						if float64(y) == roundFloat(value, 10) {
							antinodes[fmt.Sprintf("%d %d", x, y)] = struct{}{}
						}
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
