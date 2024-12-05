package main

import (
	"github.com/kiraleos/aoc-2024/util"
)

const DAY int = 4

func main() {
	input := util.FetchInput2DCharArray(DAY)

	part1 := solvePart1(input)
	part2 := solvePart2(input)

	println("Part 1: ", part1)
	println("Part 2: ", part2)
}

func solvePart1(input [][]rune) int {
	xmasCtr := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			currentRune := input[i][j]
			if currentRune == 'X' {
				if i-3 >= 0 && input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
					xmasCtr++
				}
				if i+3 < len(input) && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
					xmasCtr++
				}
				if j-3 >= 0 && input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S' {
					xmasCtr++
				}
				if j+3 < len(input[i]) && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
					xmasCtr++
				}
				if i-3 >= 0 && j-3 >= 0 && input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S' {
					xmasCtr++
				}
				if i+3 < len(input) && j+3 < len(input[i]) && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
					xmasCtr++
				}
				if i-3 >= 0 && j+3 < len(input[i]) && input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S' {
					xmasCtr++
				}
				if i+3 < len(input) && j-3 >= 0 && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
					xmasCtr++
				}
			}
		}
	}

	return xmasCtr
}

func solvePart2(input [][]rune) int {
	xmasCtr := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			currentRune := input[i][j]
			if currentRune == 'A' {
				if !(i-1 >= 0 && j-1 >= 0 && j+1 < len(input[i]) && i+1 < len(input)) {
					continue
				}
				if 
				input[i-1][j-1] == 'M' &&
				input[i-1][j+1] == 'S' &&
				input[i+1][j-1] == 'M' &&
				input[i+1][j+1] == 'S' {
					xmasCtr++
				}
				if 
				input[i-1][j-1] == 'M' &&
				input[i-1][j+1] == 'M' &&
				input[i+1][j-1] == 'S' &&
				input[i+1][j+1] == 'S' {
					xmasCtr++
				}
				if 
				input[i-1][j-1] == 'S' &&
				input[i-1][j+1] == 'M' &&
				input[i+1][j-1] == 'S' &&
				input[i+1][j+1] == 'M' {
					xmasCtr++
				}
				if 
				input[i-1][j-1] == 'S' &&
				input[i-1][j+1] == 'S' &&
				input[i+1][j-1] == 'M' &&
				input[i+1][j+1] == 'M' {
					xmasCtr++
				}
			}
		}
	}

	return xmasCtr
}
