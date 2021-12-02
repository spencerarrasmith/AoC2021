package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sample = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file() []string {
	dat, err := os.ReadFile("input.txt")
	check(err)
	datstring := string(dat)

	lines := strings.Split(datstring, "\n")

	return lines
}

func day2_part1(input []string) int {
	commands := []string{"forward", "down", "up"}
	var depth int = 0
	var position int = 0

	for _, cmd := range input {
		parsed := strings.Split(cmd, " ")

		switch parsed[0] {
		case commands[0]:
			val, _ := strconv.Atoi(parsed[1])
			position += val
		case commands[1]:
			val, _ := strconv.Atoi(parsed[1])
			depth += val
		case commands[2]:
			val, _ := strconv.Atoi(parsed[1])
			depth -= val
		}
	}
	return depth * position
}

func day2_part2(input []string) int {
	commands := []string{"forward", "down", "up"}
	var depth int = 0
	var position int = 0
	var aim int = 0

	for _, cmd := range input {
		parsed := strings.Split(cmd, " ")

		switch parsed[0] {
		case commands[0]:
			val, _ := strconv.Atoi(parsed[1])
			position += val
			depth += aim * val
		case commands[1]:
			val, _ := strconv.Atoi(parsed[1])
			aim += val
		case commands[2]:
			val, _ := strconv.Atoi(parsed[1])
			aim -= val
		}
	}
	fmt.Printf("%v\n", depth)
	fmt.Printf("%v\n", position)
	return depth * position
}

func main() {
	real := read_file()
	fmt.Printf("%v\n", day2_part1(real))
	fmt.Printf("%v\n", day2_part2(real))
}
