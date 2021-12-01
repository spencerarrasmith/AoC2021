package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sample = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file() []int {
	dat, err := os.ReadFile("input.txt")
	check(err)
	datstring := string(dat)

	raw := strings.Split(datstring, "\n")
	var lines []int
	for _, s := range raw {
		val, _ := strconv.Atoi(s)
		lines = append(lines, val)
	}
	return lines
}

func day1_part1(input []int) int {
	var answer int = 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			answer++
		}
	}

	return answer
}

func day1_part2(input []int) int {
	var answer int = 0

	for i := 3; i < len(input); i++ {
		window1 := input[i-3] + input[i-2] + input[i-1]
		window2 := input[i-2] + input[i-1] + input[i]
		if window2 > window1 {
			answer++
		}
	}

	return answer
}

func main() {
	real := read_file()
	fmt.Printf("%v\n", day1_part1(real))
	fmt.Printf("%v\n", day1_part2(real))
}
