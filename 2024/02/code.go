package main

import (
	"aoc-in-go/2024/utils"
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		values := strings.Fields(line)
		uints := make([]uint64, 0, len(values))
		for _, v := range values {
			uints = append(uints, utils.ParseUint64(v))
		}
		unsatisfiedIdx := satisfied(uints)
		if unsatisfiedIdx == -1 || (part2 && reduceUnsatisfied(uints, 0) == -1) {
			sum += 1
		}
	}

	return sum
}

func reduceUnsatisfied(uints []uint64, idx int) int {
	if len(uints) == 0 || idx >= len(uints) {
		return idx
	} else {
		newUints := utils.RemoveIndex(uints, idx)
		satisfiedIdx := satisfied(newUints)

		if satisfiedIdx == -1 {
			return satisfiedIdx
		}

		return reduceUnsatisfied(uints, idx+1)
	}
}

func satisfied(l []uint64) int {
	fulfilled := -1
	direction := ""
	length := len(l)
explore:
	for i, x := range l {
		if length == i+1 {
			break explore
		}

		switch direction {
		case "":
			if l[i+1] > x && l[i+1]-x <= 3 {
				direction = "inc"
			} else if l[i+1] < x && x-l[i+1] <= 3 {
				direction = "dec"
			} else {
				fulfilled = i
				break explore
			}
		case "inc":
			if l[i+1] > x && l[i+1]-x <= 3 {
				direction = "inc"
			} else {
				fulfilled = i
				break explore
			}
		case "dec":
			if l[i+1] < x && x-l[i+1] <= 3 {
				direction = "dec"
			} else {
				fulfilled = i
				break explore
			}
		default:
			fulfilled = i
			break explore
		}
	}

	return fulfilled
}
