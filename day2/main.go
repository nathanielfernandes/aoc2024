package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/nathanielfernandes/aoc-2024/helpers"
)

func main() {
	// part 1
	data := Must(os.ReadFile("input.txt"))

	input := string(data)
	lines := strings.Split(input, "\n")

	safeCount := 0
	for _, line := range lines {
		safe := true
		nums := strings.Split(line, " ")
		last := Must(strconv.Atoi(nums[0]))
		increasing := Must(strconv.Atoi(nums[1]))-Must(strconv.Atoi(nums[0])) >= 0

		for _, c := range nums[1:] {
			n := Must(strconv.Atoi(c))
			diff := last - n

			last = n
			if (Abs(diff) > 3) || (diff == 0) || (increasing && diff > 0) || (!increasing && diff < 0) {
				safe = false
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	fmt.Printf("Safe Reports: %d\n", safeCount)
}
