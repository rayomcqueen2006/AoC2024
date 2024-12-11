package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	NORTH direction = "North"
	EAST  direction = "South"
	SOUTH direction = "East"
	WEST  direction = "West"
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

	turns := map[direction]direction{NORTH: EAST, EAST: SOUTH, SOUTH: WEST, WEST: NORTH}

	scanner := bufio.NewScanner(file)
	grid := [][]string{}
	var x, y int = 0, 0

	// Create grid and get starting coordinates for guard
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
		possibleStartX := strings.Index(line, "^")
		if possibleStartX != -1 {
			x = possibleStartX
		}
		if x == 0 {
			y++
		}
	}

	validMove := true
	xLen := len(grid[0])
	yLen := len(grid)
	currentDirection := NORTH

	for validMove {
		grid[y][x] = "X"
		newX, newY := x, y
		moved := false
		if currentDirection == NORTH && y-1 >= 0 {
			newY = y - 1
			moved = true
		} else if currentDirection == EAST && x+1 < xLen {
			newX = x + 1
			moved = true
		} else if currentDirection == SOUTH && y+1 < yLen {
			newY = y + 1
			moved = true
		} else if currentDirection == WEST && x-1 >= 0 {
			newX = x - 1
			moved = true
		}
		if !moved {
			// guard will move off the grid
			validMove = false
			continue
		}

		// check if the guard should move or turn
		if grid[newY][newX] == "#" {
			currentDirection = turns[currentDirection]
		} else {
			x = newX
			y = newY
		}
	}

	total := 0
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			if grid[y][x] == "X" {
				total += 1
			}
		}
	}

	fmt.Println(total)
}

func partTwo() {
	// for each possible obstruction location
	// run algorithm to move guard
	// put each visited position into a map x_y_dir
	// check if the position has been visited before
	// if it has, this is a loop - increment the counter

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	turns := map[direction]direction{NORTH: EAST, EAST: SOUTH, SOUTH: WEST, WEST: NORTH}

	scanner := bufio.NewScanner(file)
	grid := [][]string{}
	var x, y int = 0, 0

	// Create grid and get starting coordinates for guard
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
		possibleStartX := strings.Index(line, "^")
		if possibleStartX != -1 {
			x = possibleStartX
		}
		if x == 0 {
			y++
		}
	}

	startX := x
	startY := y

	xLen := len(grid[0])
	yLen := len(grid)
	total := 0

	for i := 0; i < yLen; i++ {
		for j := 0; j < xLen; j++ {
			if i == startY && j == startX {
				continue
			}
			if grid[i][j] == "#" {
				continue
			}

			x = startX
			y = startY

			validMove := true
			currentDirection := NORTH
			prevValue := grid[i][j]

			// make current index an obstacle
			grid[i][j] = "#"
			positionMap := map[string]string{}

			for validMove {
				_, present := positionMap[strconv.Itoa(x)+"_"+strconv.Itoa(y)+"_"+string(currentDirection)]
				if present {
					// we are in a loop - increment counter by 1 and break out
					total += 1
					break
				} else {
					positionMap[strconv.Itoa(x)+"_"+strconv.Itoa(y)+"_"+string(currentDirection)] = ""
				}
				newX, newY := x, y
				moved := false
				if currentDirection == NORTH && y-1 >= 0 {
					newY = y - 1
					moved = true
				} else if currentDirection == EAST && x+1 < xLen {
					newX = x + 1
					moved = true
				} else if currentDirection == SOUTH && y+1 < yLen {
					newY = y + 1
					moved = true
				} else if currentDirection == WEST && x-1 >= 0 {
					newX = x - 1
					moved = true
				}
				if !moved {
					// guard will move off the grid
					validMove = false
					continue
				}

				// check if the guard should move or turn
				if grid[newY][newX] == "#" {
					currentDirection = turns[currentDirection]
				} else {
					x = newX
					y = newY
				}
			}
			grid[i][j] = prevValue
		}
	}

	fmt.Println(total)
}
