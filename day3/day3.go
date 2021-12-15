package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sample = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010"}

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

	var out []string
	for _, str := range lines {
		if len(str) != 0 {
			out = append(out, str)
		}
	}

	return out
}

func day3_part1(input []string) int64 {
	var counts = make([]int, len(input[0]))

	for _, line := range input {
		for j, char := range line {
			if char == '1' {
				counts[j]++
			}
			if char == '0' {
				counts[j]--
			}
		}
	}

	var gamma string
	var epsilon string

	for _, count := range counts {
		if count > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaval, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonval, _ := strconv.ParseInt(epsilon, 2, 64)

	return gammaval * epsilonval
}

func day3_part2(input []string) int64 {
	oxygen := input
	co2 := input
	var index int = 0

	for len(oxygen) > 1 || len(co2) > 1 {

		var oxycount int = 0
		var co2count int = 0

		for _, line := range oxygen {
			if line[index] == '1' {
				oxycount++
			}
			if line[index] == '0' {
				oxycount--
			}
		}

		for _, line := range co2 {
			if line[index] == '1' {
				co2count++
			}
			if line[index] == '0' {
				co2count--
			}
		}

		var newoxygen []string
		var newco2 []string

		var oxychar rune
		if oxycount >= 0 {
			oxychar = '1'
		} else {
			oxychar = '0'
		}

		var co2char rune
		if co2count < 0 {
			co2char = '1'
		} else {
			co2char = '0'
		}

		if len(oxygen) > 1 {
			for _, line := range oxygen {
				if rune(line[index]) == oxychar {
					newoxygen = append(newoxygen, line)
				}
			}
			oxygen = newoxygen
		}

		if len(co2) > 1 {
			for _, line := range co2 {
				if rune(line[index]) == co2char {
					newco2 = append(newco2, line)
				}
			}
			co2 = newco2
		}

		index++
	}
	oxyval, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co2val, _ := strconv.ParseInt(co2[0], 2, 64)
	return oxyval * co2val
}

func main() {
	real := read_file()
	fmt.Printf("%v\n", day3_part1(real))
	fmt.Printf("%v\n", day3_part2(real))
}
