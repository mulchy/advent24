package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// mul(X,Y)
// X,Y 3 digit positive(?) integers

type BinaryOperation struct {
	left  int
	right int
}

func calc(bs []BinaryOperation) int {
	sum := 0
	for _, v := range bs {
		sum += v.left * v.right
	}

	return sum
}

const prefix = "mul("

func findMulInstructions(input string) []BinaryOperation {
	out := make([]BinaryOperation, 0)

	possibles := strings.Split(input, prefix)
	for _, s := range possibles {
		left, right, found := strings.Cut(s, ",")
		if !found {
			continue
		}

		right, _, found = strings.Cut(right, ")")
		if !found {
			continue
		}

		l, err := strconv.Atoi(left)
		if err != nil {
			continue
		}

		r, err := strconv.Atoi(right)
		if err != nil {
			continue
		}

		out = append(out, BinaryOperation{l, r})
	}

	return out
}

func removeDisabled(s string) string {
	left, right, found := strings.Cut(s, "don't()")
	if found {
		_, rest, found := strings.Cut(right, "do()")
		if found {
			return removeDisabled(left + rest)
		}

		return left
	}

	return s
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	input = strings.ReplaceAll(input, "\n", "")

	fmt.Printf("part1: %v\n", calc(findMulInstructions(input)))
	fmt.Printf("part2: %v\n", calc(findMulInstructions(removeDisabled(input))))
}
