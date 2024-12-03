package main

import (
	"strconv"
	"strings"

	"github.com/kiraleos/aoc-2024/util"
)

const DAY int = 2

func main() {
	lines := util.FetchInputLines(DAY)

	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	println("Part 1: ", part1)
	println("Part 2: ", part2)
}

func solvePart1(lines []string) int {
	numOfSafeReports := 0
	for _, line := range lines {
		reportsStr := strings.Split(line, " ")
		report := make([]int, len(reportsStr))
		for i, reportStr := range reportsStr {
			level, err := strconv.Atoi(reportStr)
			if err != nil {
				panic("failed to parse level")
			}
			report[i] = level
		}
		if isReportValid(report) {
			numOfSafeReports++
		}
	}

	return numOfSafeReports
}

func isReportValid(report []int) bool {
	monotonicAsc := true
	monotonicDesc := true
	for i := 0; i < len(report); i++ {
		if i == 0 {
			continue
		}
		if report[i] < report[i-1] {
			monotonicAsc = false
		} else if report[i] > report[i-1] {
			monotonicDesc = false
		} else {
			monotonicAsc = false
			monotonicDesc = false
		}

		if !isDifferenceValid(report[i], report[i-1]) {
			return false
		}
	}

	isMonotonic := !(monotonicAsc && monotonicDesc) && monotonicAsc || monotonicDesc

	return isMonotonic
}

func isDifferenceValid(i int, j int) bool {
	difference := util.Abs(i - j)
	if difference < 1 || difference > 3 {
		return false
	}
	return true
}

func solvePart2(lines []string) int {
	numOfSafeReports := 0
	for _, line := range lines {
		reportsStr := strings.Split(line, " ")
		report := make([]int, len(reportsStr))
		for i, reportStr := range reportsStr {
			level, err := strconv.Atoi(reportStr)
			if err != nil {
				panic("failed to parse level")
			}
			report[i] = level
		}
		allPossibleReports := generateReportConfigurations(report)
		for _, possibleReport := range allPossibleReports {
			if isReportValid(possibleReport) {
				numOfSafeReports++
				break
			}
		}
	}

	return numOfSafeReports
}

func generateReportConfigurations(reports []int) [][]int {
	var result [][]int

	for i := range reports {
		config := append([]int{}, reports[:i]...)
		config = append(config, reports[i+1:]...)
		result = append(result, config)
	}

	return result
}
