package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"advent24/internal/utils"
)

func main() {
	input := parse()
	fmt.Println("Part 1:\n\t", part1(input))
	fmt.Println("Part 2:\n\t", part2(input))
}

func parse() [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		ints := make([]int, len(fields))

		for i, v := range fields {
			converted, err := strconv.Atoi(v)
			if err != nil {
				panic(fmt.Sprintf("Failed to parse this string %s as an int", fields[i]))
			}

			ints[i] = converted
		}

		lines = append(lines, ints)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func in_range(v1 int, v2 int) bool {
	diff := utils.Abs(v1 - v2)
	return diff >= 1 && diff <= 3
}

func safe(report []int) bool {
	increasing := true
	decreasing := true
	valid := false

	for i := range report {
		if i == 0 {
			continue
		}

		v1 := report[i-1]
		v2 := report[i]

		if v1 < v2 {
			decreasing = false
		}

		if v1 > v2 {
			increasing = false
		}

		valid = in_range(v1, v2)

		if !valid {
			break
		}
	}
	monotonic := increasing || decreasing

	return monotonic && valid
}

func part1(input [][]int) int {
	safeReports := 0

	for _, report := range input {
		if safe(report) {
			safeReports++
		}
	}

	return safeReports
}

func part2(input [][]int) int {
	safeReports := 0

	for _, report := range input {
		for i := range report {
			// probably not the most efficient way to do this
			newReport := make([]int, 0)
			newReport = append(newReport, report[:i]...)
			newReport = append(newReport, report[i+1:]...)

			if safe(newReport) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}
