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
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		distances := strings.Split(line, "   ")
		distance1, err := strconv.Atoi(distances[0])
		distance2, err := strconv.Atoi(distances[1])
		if err != nil {
			fmt.Println("Error converting to integer: " + err.Error())
		}

		list1 = append(list1, distance1)
		list2 = append(list2, distance2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	total := 0
	for i := 0; i < len(list1); i++ {
		total += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(total)
}

func partTwo() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
	}

	list1 := []int{}
	mapOfList2 := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		distances := strings.Split(line, "   ")
		distance1, err := strconv.Atoi(distances[0])
		distance2, err := strconv.Atoi(distances[1])
		if err != nil {
			fmt.Println("Error converting to integer: " + err.Error())
		}

		list1 = append(list1, distance1)

		count, present := mapOfList2[distance2]
		if present {
			mapOfList2[distance2] = count + 1
		} else {
			mapOfList2[distance2] = 1
		}
	}

	// get total by finding how many times each number in list 1 is in list 2
	total := 0
	for _, distance := range list1 {
		score := distance * mapOfList2[distance]
		total += score
	}

	fmt.Println(total)
}
