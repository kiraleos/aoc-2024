package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/kiraleos/aoc-2024/util"
)

const DAY int = 1

func main() {
	lines := util.FetchInputLines(DAY)
	
	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	println("Part 1: ", part1)
	println("Part 2: ", part2)
}

func solvePart1(lines []string) int {
	firstColumn, secondColumn := splitToSortedColumns(lines)
	totalDistance := calculateSumOfDistanceBetweenColumns(firstColumn, secondColumn)
	return totalDistance
}

func solvePart2(lines []string) int {
	firstColumn, secondColumn := splitToSortedColumns(lines)
	frequencies := calculateFrequencies(firstColumn, secondColumn)
	similarityScore := calculateSimilarityScore(firstColumn, frequencies)
	return similarityScore
}

func calculateSimilarityScore(firstColumn []int, frequencies map[int]int) int {
	similarityScore := 0
	for i := 0; i < len(firstColumn); i++ {
		similarityScore += firstColumn[i] * frequencies[firstColumn[i]]
	}
	return similarityScore
}

func calculateFrequencies(firstColumn []int, secondColumn []int) map[int]int {
	frequencies := make(map[int]int)
	for i := 0; i < len(firstColumn); i++ {
		for j := 0; j < len(secondColumn); j++ {
			if firstColumn[i] == secondColumn[j] {
				frequencies[firstColumn[i]]++
			}
		}
	}
	return frequencies
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