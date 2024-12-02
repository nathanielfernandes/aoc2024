package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	. "github.com/nathanielfernandes/aoc-2024/helpers"
)

func main() {
	// part 1
	data := Must(os.ReadFile("input.txt"))

	input := string(data)
	lines := strings.Split(input, "\n")

	n := len(lines)
	left := make([]int, 0, n)
	right := make([]int, 0, n)

	for _, line := range lines {
		pair := strings.Split(line, "   ")
		left = append(left, Must(strconv.Atoi(strings.TrimSpace(pair[0]))))
		right = append(right, Must(strconv.Atoi(strings.TrimSpace(pair[1]))))
	}

	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0
	for i := range n {
		totalDistance += Abs(left[i] - right[i])
	}

	// first answer
	fmt.Printf("First Answer: %d\n", totalDistance)

	// part 2
	counts := make(map[int]int, n)
	for _, v := range right {
		if prev, ok := counts[v]; ok {
			counts[v] = prev + 1
		} else {
			counts[v] = 1
		}
	}

	similarityScore := 0
	for _, v := range left {
		if count, ok := counts[v]; ok {
			similarityScore += v * count
		}
	}

	// second answer
	fmt.Printf("Second Answer: %d\n", similarityScore)
}
