package main

import (
	"fmt"
	"strings"

	"gitea.kandjdev.net/KandJDev/aocutils/files"
)

type position struct {
	x int
	y int
}

type plot struct {
	position position
	value    string
	north    string
	east     string
	south    string
	west     string
	visited  bool
	inRegion bool
}

func main() {
	partOne()
}

func partOne() {
	total := 0
	garden := [][]string{}
	files.ScanFile("input.txt", func(line string) {
		row := strings.Split(line, "")
		garden = append(garden, row)
	})

	// build map of plots
	plots := map[position]plot{}
	for y := range len(garden) {
		for x := range len(garden[y]) {
			var n, s, e, w string
			if y-1 >= 0 {
				n = garden[y-1][x]
			}
			if x+1 < len(garden[0]) {
				e = garden[y][x+1]
			}
			if y+1 < len(garden) {
				s = garden[y+1][x]
			}
			if x-1 >= 0 {
				w = garden[y][x-1]
			}
			plots[position{x: x, y: y}] = plot{
				position: position{x: x, y: y},
				value:    garden[y][x],
				north:    n,
				east:     e,
				south:    s,
				west:     w,
			}
		}
	}

	// traverse garden and collect regions
	for y := range len(garden) {
		for x := range len(garden[y]) {
			if plots[position{x: x, y: y}].inRegion {
				continue
			}
			// we haven't explored this region yet
			region := map[plot]struct{}{}
			visited := map[position]bool{}
			dfs(position{x: x, y: y}, plots[position{x: x, y: y}].value, plots, visited, region)
			// reset visited
			for _, p := range plots {
				plots[p.position] = plot{
					position: p.position,
					value:    p.value,
					visited:  false,
					north:    p.north,
					east:     p.east,
					south:    p.south,
					west:     p.west,
					inRegion: p.inRegion,
				}
			}
			area := len(region)
			perimeter := 0
			// set in region
			for p := range region {
				if p.north != p.value {
					perimeter++
				}
				if p.east != p.value {
					perimeter++
				}
				if p.south != p.value {
					perimeter++
				}
				if p.west != p.value {
					perimeter++
				}
				plots[p.position] = plot{
					position: p.position,
					value:    p.value,
					visited:  p.visited,
					north:    p.north,
					east:     p.east,
					south:    p.south,
					west:     p.west,
					inRegion: true,
				}
			}
			total += (area * perimeter)
		}
	}

	fmt.Println(total)
}

func dfs(
	current position,
	targetVal string,
	plots map[position]plot,
	visited map[position]bool,
	region map[plot]struct{},
) {
	if visited[current] {
		return
	}
	if plots[current].value != targetVal {
		return
	}

	visited[current] = true
	region[plots[current]] = struct{}{}

	dirs := []position{
		{current.x, current.y - 1},
		{current.x + 1, current.y},
		{current.x, current.y + 1},
		{current.x - 1, current.y},
	}

	for _, next := range dirs {
		if _, ok := plots[next]; ok {
			dfs(next, targetVal, plots, visited, region)
		}
	}
}
