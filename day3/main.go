package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	. "github.com/nathanielfernandes/aoc-2024/helpers"
)

func parseMulInner(inp string, i int) (int, int) {
	// string buffers
	lhs := strings.Builder{}
	rhs := strings.Builder{}
	lhs.Grow(3)
	rhs.Grow(3)

	// if we have seen a comma
	comma := false
	// only look ahead 7 chars, because the numbers are 1-3 digits + `)`
	j := i
	for j < i+7 {
		n := inp[j]
		if n == ',' {
			if comma {
				j++
				break
			}
			comma = true
		} else if unicode.IsDigit(rune(n)) {
			if comma {
				rhs.WriteByte(n)
			} else {
				lhs.WriteByte(n)
			}
		} else {
			break
		}
		j++
	}

	// should be left with a `)` or it is corrupted
	if inp[j] == ')' && comma && lhs.Len() > 0 && rhs.Len() > 0 {
		return Must(strconv.Atoi(lhs.String())) * Must(strconv.Atoi(rhs.String())), j
	}

	return 0, j
}

func part1(inp string) int {
	total := 0
	n := len(inp)
	i := 0
	for i < n-4 {
		// if we see a `mul(`
		if inp[i] == 'm' && inp[i+1] == 'u' && inp[i+2] == 'l' && inp[i+3] == '(' {
			// eat the chars
			i += 4

			var t int
			t, i = parseMulInner(inp, i)
			total += t
		}
		i++
	}

	return total
}

func part2(inp string) int {
	do := true
	total := 0

	n := len(inp)
	i := 0
	for i < n-7 {
		// if we see a `mul(`
		if do && inp[i] == 'm' && inp[i+1] == 'u' && inp[i+2] == 'l' && inp[i+3] == '(' {
			// eat the chars
			i += 4
			var t int
			t, i = parseMulInner(inp, i)
			total += t

			// if we see `do`
		} else if inp[i] == 'd' && inp[i+1] == 'o' {
			i += 2

			// is this `don't`
			enable := true
			if inp[i] == 'n' && inp[i+1] == '\'' && inp[i+2] == 't' {
				i += 3
				enable = false
			}

			if inp[i] == '(' && inp[i+1] == ')' {
				i += 2
				do = enable
			}
		} else {
			i++
		}
	}

	return total
}

func main() {
	// part 1
	data := Must(os.ReadFile("input.txt"))
	inp := string(data)

	fmt.Printf("Total: %d\n", part1(inp))
	fmt.Printf("Total P2: %d\n", part2(inp))
}
