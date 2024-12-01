package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/kiraleos/aoc-2024/util"
)

const DAY int = 1

func main() {
	lines := fetchInputLines(DAY)
	
	firstColumn, secondColumn := splitToSortedColumns(lines)
	totalDistance := calculateSumOfDistanceBetweenColumns(firstColumn, secondColumn)

	println(totalDistance)
}

func fetchInputLines(day int) []string {
	input, err := util.FetchInput(day)
	if err != nil {
		panic(err)
	}

	lines, err := util.SplitInputToLines(input)
	if err != nil {
		panic(err)
	}
	return lines
}

func calculateSumOfDistanceBetweenColumns(firstColumn []int, secondColumn []int) int {
	totalDistance := 0
	for i := 0; i < len(firstColumn); i++ {
		distance := Abs(firstColumn[i] - secondColumn[i])
		totalDistance += distance
	}
	return totalDistance
}

func splitToSortedColumns(lines []string) ([]int, []int) {
	var firstColumn []int
	var secondColumn []int

	for _, line := range lines {
		line := strings.Split(line, "   ")
		num1, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		firstColumn = append(firstColumn, num1)
		secondColumn = append(secondColumn, num2)
	}

	sort.Slice(firstColumn, func(i, j int) bool {
		return firstColumn[i] < firstColumn[j]
	})
	sort.Slice(secondColumn, func(i, j int) bool {
		return secondColumn[i] < secondColumn[j]
	})

	return firstColumn, secondColumn
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}