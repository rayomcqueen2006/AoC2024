package main

import (
	"fmt"
	"strconv"
	"strings"

	"gitea.kandjdev.net/KandJDev/aocutils/files"
)

type position struct {
	x int
	y int
}

type node struct {
	children []position
	value    string
	position position
}

var total int = 0

func main() {
	partOne()
	partTwo()
}

func partOne() {
	topoMap := [][]string{}
	files.ScanFile("input.txt", func(line string) {
		row := strings.Split(line, "")
		topoMap = append(topoMap, row)
	})

	overallTotal := 0
	// for each trailhead, assemble a tree with nodes one apart and increasing
	for y := range topoMap {
		for x := range topoMap[y] {
			if topoMap[y][x] == "0" {
				trailTree := generateTrailTree(topoMap, x, y)

				// Calculate how many nodes with no children have a value of 9 (these are the tops of the trails)
				trailheadTotal := 0
				for _, node := range trailTree {
					if len(node.children) == 0 && node.value == "9" {
						trailheadTotal++
					}
				}
				fmt.Println("trailhead position", x, y, "- total: ", trailheadTotal)
				overallTotal += trailheadTotal
			}
		}
	}
	fmt.Println("overall total: ", overallTotal)
}

func partTwo() {
	topoMap := [][]string{}
	files.ScanFile("input.txt", func(line string) {
		row := strings.Split(line, "")
		topoMap = append(topoMap, row)
	})

	// for each trailhead, assemble a tree with nodes one apart and increasing
	for y := range topoMap {
		for x := range topoMap[y] {
			if topoMap[y][x] == "0" {
				trailTree := generateTrailTree(topoMap, x, y)

				// Do a DFS to find all the paths
				stack := make([]position, 0)
				dfs(position{x: x, y: y}, "9", stack, trailTree)
			}
		}
	}
	fmt.Println("overall total: ", total)
}

func dfs(current position, endVal string, stack []position, trailTree map[position]node) bool {
	stack = append(stack, current)
	if trailTree[current].value == endVal {
		return true
	}
	for _, child := range trailTree[current].children {
		if dfs(child, endVal, stack, trailTree) {
			total += 1
		}
	}
	stack = stack[:len(stack)-1]
	return false
}

func generateTrailTree(topoMap [][]string, x, y int) map[position]node {
	trailTree := map[position]node{}

	// add the root node as the trailhead
	trailTree[position{x, y}] = node{children: []position{}, value: "0", position: position{x, y}}

	nodesToAdd := map[position]node{}
	start := true
	// when no more nodes are added we are at the bottom layer of the tree
	for len(nodesToAdd) != 0 || start {
		// reset the new nodes to be added to this layer of the tree
		nodesToAdd = map[position]node{}
		for pos, trailNode := range trailTree {
			if len(trailNode.children) == 0 {
				// We haven't yet explored this node so let's take a look at its neighbours
				childrenToAddToCurrentNode := []position{}
				// Top
				if pos.y-1 >= 0 {
					checkNode(trailTree, topoMap, trailNode, pos.x, pos.y-1, nodesToAdd, &childrenToAddToCurrentNode)
				}
				// Right
				if pos.x+1 < len(topoMap[0]) {
					checkNode(trailTree, topoMap, trailNode, pos.x+1, pos.y, nodesToAdd, &childrenToAddToCurrentNode)
				}
				// Bottom
				if pos.y+1 < len(topoMap) {
					checkNode(trailTree, topoMap, trailNode, pos.x, pos.y+1, nodesToAdd, &childrenToAddToCurrentNode)
				}
				// Left
				if pos.x-1 >= 0 {
					checkNode(trailTree, topoMap, trailNode, pos.x-1, pos.y, nodesToAdd, &childrenToAddToCurrentNode)
				}
				if len(childrenToAddToCurrentNode) > 0 {
					nodesToAdd[trailNode.position] = node{
						value:    trailNode.value,
						position: trailNode.position,
						children: childrenToAddToCurrentNode,
					}
				}
			}
		}
		for _, node := range nodesToAdd {
			trailTree[node.position] = node
		}
		start = false
	}
	return trailTree
}

func checkNode(trailTree map[position]node, topoMap [][]string, currentNode node, x, y int, nodesToAdd map[position]node, childrenToAdd *[]position) {
	currentNodeVal, _ := strconv.Atoi(currentNode.value)
	nodeToCheck := topoMap[y][x]
	nodeToCheckVal, _ := strconv.Atoi(nodeToCheck)
	if currentNodeVal+1 == nodeToCheckVal {
		newPos := position{x: x, y: y}
		// Add the new node to the tree if it doesn't already exist
		if _, ok := trailTree[newPos]; !ok {
			nodesToAdd[newPos] = node{position: position{x, y}, children: []position{}, value: nodeToCheck}
		}
		// Make sure to add the node to the current node's children
		*childrenToAdd = append(*childrenToAdd, newPos)
	}
}
