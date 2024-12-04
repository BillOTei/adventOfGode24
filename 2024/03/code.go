package main

import (
	"aoc-in-go/2024/utils"
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := uint64(0)
	reNb := regexp.MustCompile("[0-9]+")

	if part2 {
		r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|(d)(o)(?:n't)?\(\)`)
		allMatch := r.FindAllString(input, -1)
		enabled := true
		for _, match := range allMatch {
			switch match {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					numbers := reNb.FindAllString(match, -1)
					sum += utils.ParseUint64(numbers[0]) * utils.ParseUint64(numbers[1])
				}
			}
		}

		return sum
	}

	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	allMatch := r.FindAllString(input, -1)
	for _, match := range allMatch {
		numbers := reNb.FindAllString(match, -1)
		sum += utils.ParseUint64(numbers[0]) * utils.ParseUint64(numbers[1])
	}

	return sum
}
