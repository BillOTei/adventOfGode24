package main

import (
	"aoc-in-go/2024/utils"
	"github.com/jpillora/puzzler/harness/aoc"
	"slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := uint64(0)
	stringSlice := strings.Split(input, "\n\n")
	rawMap := strings.Split(stringSlice[0], "\n")
	rawPages := strings.Split(stringSlice[1], "\n")
	mPred := make(map[uint64][]uint64)
	mSucc := make(map[uint64][]uint64)
	for _, couple := range rawMap {
		pageNumbers := strings.Split(couple, "|")
		k := utils.ParseUint64(pageNumbers[1])
		v := utils.ParseUint64(pageNumbers[0])
		mPred[k] = append(mPred[k], v)
		mSucc[v] = append(mSucc[v], k)
	}
	for _, update := range rawPages {
		pages := strings.Split(update, ",")
		valid := true
		var buff []uint64
	pageLoop:
		for _, p := range pages {
			page := utils.ParseUint64(p)
			predecessors := mPred[page]
			successors := mSucc[page]
			if len(buff) == 0 {
				buff = append(buff, page)
			} else {
				for _, succ := range successors {
					if slices.Contains(pages, strconv.FormatUint(succ, 10)) && slices.Contains(buff, succ) {
						valid = false
						break pageLoop
					}
				}
				for _, pred := range predecessors {
					if slices.Contains(pages, strconv.FormatUint(pred, 10)) && !slices.Contains(buff, pred) {
						valid = false
						break pageLoop
					}
				}
				buff = append(buff, page)
			}
		}
		if valid {
			l := len(pages)
			sum += utils.ParseUint64(pages[l/2])
		}
	}

	if part2 {
		return "not implemented"
	}

	return sum
}
