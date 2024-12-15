package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	limit := len(lines)
	visited := move(
		&lines,
		limit,
		point{
			x:         60,
			y:         36,
			direction: "^",
		},
		make(map[p]bool),
	)

	if part2 {
		return "not implemented"
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

type point struct {
	x         int
	y         int
	direction string
}

type p struct {
	x int
	y int
}
