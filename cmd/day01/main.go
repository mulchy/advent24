package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"advent24/internal/utils"
)

func main() {
	list1, list2 := parse()
	fmt.Println("Part 1:\n\t", part1(list1, list2))
	fmt.Println("Part 2:\n\t", part2(list1, list2))
}

func parse() ([]int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		one, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(fmt.Sprintf("Failed to parse int %s", fields[0]))
		}
		list1 = append(list1, one)

		two, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(fmt.Sprintf("Failed to parse int %s", fields[1]))
		}
		list2 = append(list2, two)
	}
	return list1, list2
}

func part1(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for i, one := range list1 {
		two := list2[i]
		diff := utils.Abs(one - two)

		sum += diff
	}

	return sum
}

func part2(list1 []int, list2 []int) int {
	similarityScore := 0

	for _, id := range list1 {

		count := 0
		for _, other := range list2 {
			if id == other {
				count++
			}
		}

		similarityScore += id * count
	}

	return similarityScore
}
