package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	sum := 0
	//xmas := regexp.MustCompile("XMAS")
	//samx := regexp.MustCompile("SAMX")
	//for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
	//	allXmas := xmas.FindAllString(line, -1)
	//	allSamx := samx.FindAllString(line, -1)
	//	fmt.Println(allXmas, allSamx)
	//	sum += len(allXmas) + len(allSamx)
	//}

	lines := strings.Split(strings.TrimSpace(input), "\n")
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

	if part2 {
		return runPart2(input)
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

func checkIfElementsPart2(inputOne string, inputTwo string) bool {
	//Check if the inputs provided creates a X-MAS one way or the other
	if (inputOne == "MAS" || inputOne == "SAM") && (inputTwo == "MAS" || inputTwo == "SAM") {
		return true
	}
	return false
}
func runPart2(input string) int {
	//Get  rows of data
	rows := strings.Split(input, "\r\n")
	var count = 0
	//Loop Through Rows
	for i := 0; i <= (len(rows) - 1); i++ {
		if (len(rows) - 1) <= i {
			break
		}
		//Skip first and last Row as cannot make an X-MAS with no rows above or below
		if i != 0 && i <= (len(rows)-1) {
			//Loop through value of Rows
			for z := 0; z < len(rows[i]); z++ {
				//if Value is "A" check the correct values the row above and below and Check Elements to see if they make and X-MAS
				if string(rows[i][z]) == "A" && len(rows[i][:z]) >= 1 && len(rows[i][z:]) > 1 {
					if checkIfElementsPart2(string(string(rows[i-1][z-1])+string(rows[i][z])+string(rows[i+1][z+1])), string(string(rows[i+1][z-1])+string(rows[i][z])+string(rows[i-1][z+1]))) {
						count++
					}
				}
			}
		}
	}

	return count
}
