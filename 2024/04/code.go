package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	//xmas := regexp.MustCompile("XMAS")
	//samx := regexp.MustCompile("SAMX")
	//for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
	//	allXmas := xmas.FindAllString(line, -1)
	//	allSamx := samx.FindAllString(line, -1)
	//	fmt.Println(allXmas, allSamx)
	//	sum += len(allXmas) + len(allSamx)
	//}

	if part2 {
		//mask1 := make([][]string, 3)
		//mask2 := make([][]string, 3)
		//mask3 := make([][]string, 3)
		//mask4 := make([][]string, 3)
		//for y := 0; y < 3; y++ {
		//	mask1[y] = make([]string, 3)
		//	mask2[y] = make([]string, 3)
		//	mask3[y] = make([]string, 3)
		//	mask4[y] = make([]string, 3)
		//
		//	for x := 0; x < 3; x++ {
		//		switch x {
		//		case 0:
		//			switch y {
		//			case 0:
		//				mask1[y][x] = "M"
		//				mask2[y][x] = "S"
		//				mask3[y][x] = "M"
		//				mask4[y][x] = "S"
		//			case 1:
		//				mask1[y][x] = "."
		//				mask2[y][x] = "."
		//				mask3[y][x] = "."
		//				mask4[y][x] = "."
		//			case 2:
		//				mask1[y][x] = "S"
		//				mask2[y][x] = "M"
		//				mask3[y][x] = "M"
		//				mask4[y][x] = "A"
		//			}
		//		case 1:
		//			switch y {
		//			case 0:
		//				mask1[y][x] = "."
		//				mask2[y][x] = "."
		//				mask3[y][x] = "."
		//				mask4[y][x] = "."
		//			case 1:
		//				mask1[y][x] = "A"
		//				mask2[y][x] = "A"
		//				mask3[y][x] = "A"
		//				mask4[y][x] = "A"
		//			case 2:
		//				mask1[y][x] = "."
		//				mask2[y][x] = "."
		//				mask3[y][x] = "."
		//				mask4[y][x] = "."
		//			}
		//		case 2:
		//			switch y {
		//			case 0:
		//				mask1[y][x] = "M"
		//				mask2[y][x] = "S"
		//				mask3[y][x] = "S"
		//				mask4[y][x] = "M"
		//			case 1:
		//				mask1[y][x] = "."
		//				mask2[y][x] = "."
		//				mask3[y][x] = "."
		//				mask4[y][x] = "."
		//			case 2:
		//				mask1[y][x] = "S"
		//				mask2[y][x] = "M"
		//				mask3[y][x] = "S"
		//				mask4[y][x] = "M"
		//			}
		//		}
		//	}
		//}

		length := len(lines)
		for y, line := range lines {
			runes := []rune(line)
			for x, c := range runes {
				if string(c) == "A" && x != -1 && x > 0 && x < length-1 && y > 0 && y < length-1 {
					up := lines[y-1][x-1 : x+2]
					down := lines[y+1][x-1 : x+2]

					//fmt.Println(lines[y-1], up, lines[y+1], down)

					m1Up, _ := regexp.MatchString("M.M", up)
					m1Down, _ := regexp.MatchString("S.S", down)
					m2Up, _ := regexp.MatchString("M.S", up)
					m2Down, _ := regexp.MatchString("M.S", down)
					m3Up, _ := regexp.MatchString("S.S", up)
					m3Down, _ := regexp.MatchString("M.M", down)
					m4Up, _ := regexp.MatchString("S.M", up)
					m4Down, _ := regexp.MatchString("S.M", down)
					if (m1Up && m1Down) || (m2Up && m2Down) || (m3Up && m3Down) || (m4Up && m4Down) {
						sum += 1
					}
				}
			}

		}

		return sum
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			neighbours := newNeighbours(j, i)
			for _, n := range neighbours.all() {
				if n.x < 0 || n.y < 0 || n.y >= len(lines) || n.x >= len(lines[i]) {
					continue
				}

				if lines[i][j] == 'X' {
					if checkNeighbours(n, "MAS", lines) {
						sum++
					}
				}
			}
		}
	}

	return sum
}

func checkNeighbours(n neighbour, word string, lines []string) bool {
	if len(word) == 0 {
		return true
	}

	if n.x < 0 || n.y < 0 || n.y >= len(lines) || n.x >= len(lines[n.y]) {
		return false
	}

	if lines[n.y][n.x] != word[0] {
		return false
	}

	return checkNeighbours(newNeighbour(n.direction, n.x, n.y), word[1:], lines)
}

type neighbour struct {
	x         int
	y         int
	direction int
}

func newNeighbour(direction, x, y int) neighbour {
	switch direction {
	case N:
		return neighbour{x, y + 1, direction}
	case NE:
		return neighbour{x + 1, y + 1, direction}
	case E:
		return neighbour{x + 1, y, direction}
	case SE:
		return neighbour{x + 1, y - 1, direction}
	case S:
		return neighbour{x, y - 1, direction}
	case SW:
		return neighbour{x - 1, y - 1, direction}
	case W:
		return neighbour{x - 1, y, direction}
	case NW:
		return neighbour{x - 1, y + 1, direction}
	default:
		return neighbour{}
	}
}

type neighbours struct {
	N  neighbour
	NE neighbour
	E  neighbour
	SE neighbour
	S  neighbour
	SW neighbour
	W  neighbour
	NW neighbour
}

func newNeighbours(x, y int) neighbours {
	return neighbours{
		newNeighbour(N, x, y),
		newNeighbour(NE, x, y),
		newNeighbour(E, x, y),
		newNeighbour(SE, x, y),
		newNeighbour(S, x, y),
		newNeighbour(SW, x, y),
		newNeighbour(W, x, y),
		newNeighbour(NW, x, y),
	}
}

func (n neighbours) all() [8]neighbour {
	return [8]neighbour{n.N, n.NE, n.E, n.SE, n.S, n.SW, n.W, n.NW}
}

const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)
