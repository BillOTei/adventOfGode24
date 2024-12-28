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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	limit := len(lines)
	start := point{
		x:         60,
		y:         36,
		direction: "^",
	}
	visited := move(
		&lines,
		limit,
		start,
		make(map[p]bool),
	)

	if part2 {
		blocks := make(map[p]bool)
		for toCheck := range visited {
			newLines := make([]string, limit)
			copy(newLines, lines)
			l := newLines[toCheck.y]
			newLines[toCheck.y] = utils.ReplaceAtIndex(l, '#', toCheck.x)
			if move2(&newLines, limit, start, make(map[point]bool)) {
				blocks[toCheck] = true
			}
		}

		return len(blocks)
	}

	return len(visited)
}

func move(m *[]string, limit int, currP point, visited map[p]bool) map[p]bool {

	//fmt.Println(currP, visited)

	visited[p{
		x: currP.x,
		y: currP.y,
	}] = true
	switch currP.direction {
	case "^":
		if currP.y-1 < 0 {
			return visited
		}
		l := (*m)[currP.y-1]
		next := string([]rune(l)[currP.x])
		if next == "#" {
			return move(
				m,
				limit,
				point{
					x:         currP.x + 1,
					y:         currP.y,
					direction: ">",
				},
				visited,
			)
		} else {
			return move(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y - 1,
					direction: "^",
				},
				visited,
			)
		}

	case ">":
		if currP.x+1 >= limit {
			return visited
		}
		l := (*m)[currP.y]
		next := string([]rune(l)[currP.x+1])
		if next == "#" {
			return move(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y + 1,
					direction: "v",
				},
				visited,
			)
		} else {
			return move(
				m,
				limit,
				point{
					x:         currP.x + 1,
					y:         currP.y,
					direction: ">",
				},
				visited,
			)
		}
	case "v":
		if currP.y+1 >= limit {
			return visited
		}
		l := (*m)[currP.y+1]
		next := string([]rune(l)[currP.x])
		if next == "#" {
			return move(
				m,
				limit,
				point{
					x:         currP.x - 1,
					y:         currP.y,
					direction: "<",
				},
				visited,
			)
		} else {
			return move(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y + 1,
					direction: "v",
				},
				visited,
			)
		}
	case "<":
		if currP.x-1 < 0 {
			return visited
		}
		l := (*m)[currP.y]
		next := string([]rune(l)[currP.x-1])
		if next == "#" {
			return move(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y - 1,
					direction: "^",
				},
				visited,
			)
		} else {
			return move(
				m,
				limit,
				point{
					x:         currP.x - 1,
					y:         currP.y,
					direction: "<",
				},
				visited,
			)
		}
	}

	return visited
}

func move2(m *[]string, limit int, currP point, visited map[point]bool) bool {

	//fmt.Println(currP, visited)

	if _, ok := visited[currP]; ok {
		return true
	} else {
		visited[currP] = true
	}

	switch currP.direction {
	case "^":
		if currP.y-1 < 0 {
			return false
		}
		l := (*m)[currP.y-1]
		next := string([]rune(l)[currP.x])
		if next == "#" {
			return move2(
				m,
				limit,
				point{
					x:         currP.x + 1,
					y:         currP.y,
					direction: ">",
				},
				visited,
			)
		} else {
			return move2(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y - 1,
					direction: "^",
				},
				visited,
			)
		}

	case ">":
		if currP.x+1 >= limit {
			return false
		}
		l := (*m)[currP.y]
		next := string([]rune(l)[currP.x+1])
		if next == "#" {
			return move2(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y + 1,
					direction: "v",
				},
				visited,
			)
		} else {
			return move2(
				m,
				limit,
				point{
					x:         currP.x + 1,
					y:         currP.y,
					direction: ">",
				},
				visited,
			)
		}
	case "v":
		if currP.y+1 >= limit {
			return false
		}
		l := (*m)[currP.y+1]
		next := string([]rune(l)[currP.x])
		if next == "#" {
			return move2(
				m,
				limit,
				point{
					x:         currP.x - 1,
					y:         currP.y,
					direction: "<",
				},
				visited,
			)
		} else {
			return move2(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y + 1,
					direction: "v",
				},
				visited,
			)
		}
	case "<":
		if currP.x-1 < 0 {
			return false
		}
		l := (*m)[currP.y]
		next := string([]rune(l)[currP.x-1])
		if next == "#" {
			return move2(
				m,
				limit,
				point{
					x:         currP.x,
					y:         currP.y - 1,
					direction: "^",
				},
				visited,
			)
		} else {
			return move2(
				m,
				limit,
				point{
					x:         currP.x - 1,
					y:         currP.y,
					direction: "<",
				},
				visited,
			)
		}
	}

	return false
}

type point struct {
	x         int
	y         int
	direction string
}

type p struct {
	x int
	y int
}
