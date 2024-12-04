package main

import (
	"aoc-in-go/2024/utils"
	"github.com/jpillora/puzzler/harness/aoc"
	"sort"
	"strings"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	sum := uint64(0)
	var left []uint64
	var right []uint64
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		values := strings.Fields(line)
		left = append(left, utils.ParseUint64(values[0]))
		right = append(right, utils.ParseUint64(values[1]))
	}

	if part2 {
		m := make(map[uint64]uint64)
		for _, v := range right {
			m[v]++
		}
		for _, v := range left {
			sum += v * m[v]
		}

		return sum
	}

	sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })
	var length = len(left)
	for i := 0; i < length; i++ {
		sum += diff(left[i], right[i])
	}

	return sum
}

func diff(a, b uint64) uint64 {
	if a < b {
		return b - a
	}
	return a - b
}
