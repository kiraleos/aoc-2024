package main

import (
	"strconv"
	"strings"

	"github.com/kiraleos/aoc-2024/util"
)

const DAY int = 5

func main() {
	input := util.FetchInput(DAY)

	pageOrder, updates := parseInput(input)

	part1 := solvePart1(pageOrder, updates)
	part2 := solvePart2(pageOrder, updates)

	println("Part 1: ", part1)
	println("Part 2: ", part2)
}

func parseInput(input string) (pageOrder [][]int, updates [][]int) {
	parts := strings.Split(input, "\n\n")
	pageOrderStr := strings.Split(parts[0], "\n")
	pageOrder = make([][]int, len(pageOrderStr))

	for i, line := range pageOrderStr {
		pageOrder[i] = make([]int, 2)
		pages := strings.Split(line, "|")
		pageOrder[i][0], _ = strconv.Atoi(pages[0])
		pageOrder[i][1], _ = strconv.Atoi(pages[1])
	}

	updatesStr := strings.Split(parts[1], "\n")
	updates = make([][]int, len(updatesStr))

	for i, line := range updatesStr {
		update := strings.Split(line, ",")
		updates[i] = make([]int, len(update))
		for j, num := range update {
			updates[i][j], _ = strconv.Atoi(num)
		}
	}

	return pageOrder, updates[:len(updates)-1]
}

func solvePart1(pageOrder [][]int, updates [][]int) int {
	pageOrders := make(map[int][]int)
	for _, pages := range pageOrder {
		pageOrders[pages[0]] = append(pageOrders[pages[0]], pages[1])
	}
	correctUpdates := make([]bool, len(updates))
	for i := range updates {
		correctUpdates[i] = true
	}

	for _, update := range updates {
		for i, currPageInUpdate := range update {
			if i == 0 {
				continue
			}
			pagesBefore := update[:i]
			for page := range pagesBefore {
				for _, pageOrder := range pageOrders[currPageInUpdate] {
					if pageOrder == page {
						correctUpdates[i] = false
					}
				}
			}
		}
	}

	middlePageSum := 0

	for i, update := range updates {
		if correctUpdates[i] {
			if len(update) % 2 == 0 {
				middlePageSum += update[len(update)/2]
			} else {
				middlePageSum += update[len(update)/2+1]
			}
		}
	}

	return middlePageSum
}

func solvePart2(pageOrder [][]int, updates [][]int) int {
	return 2
}
