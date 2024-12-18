package main

import (
	"regexp"
	"strconv"

	"github.com/kiraleos/aoc-2024/util"
)

const DAY int = 3

func main() {
	memory := util.FetchInput(DAY)

	part1 := solvePart1(memory)
	part2 := solvePart2(memory)

	println("Part 1: ", part1)
	println("Part 2: ", part2)
}

func solvePart1(memory string) int {
	pattern := `mul\((\d+),(\d+)\)`
	regex, _ := regexp.Compile(pattern)
	muls := regex.FindAllStringSubmatch(memory, -1)

	mulSum := 0
	for _, match := range muls {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		mulSum += num1 * num2
	}

	return mulSum
}

func solvePart2(memory string) int {
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`
	regex, _ := regexp.Compile(pattern)
	matches := regex.FindAllStringSubmatch(memory, -1)

	doMuls := true
	mulSum := 0
	for _, match := range matches {
		if match[0] == "do()" {
			doMuls = true
		} else if match[0] == "don't()" {
			doMuls = false
		}

		if doMuls {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			mulSum += num1 * num2
		}
	}

	return mulSum
}
