package main

import "github.com/kiraleos/aoc-2024/util"

func main() {
	day := 1
	input, err := util.FetchInput(day)
	if err != nil {
		panic(err)
	}
	println(input)
}
