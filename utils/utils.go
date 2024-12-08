package utils

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Day interface {
	Id() int
	LoadFile(isTest bool) []string
	Part1(lines []string) (int, error)
	Part2(lines []string) (int, error)
}

func LoadFile(isTest bool, dayId int) []string {
	var fileName string

	if isTest {
		fileName = "test"
	} else {
		fileName = "input"
	}

	dirName := fmt.Sprintf("day%d", dayId)

	file, err := os.ReadFile(fmt.Sprintf("%s/%s.txt", dirName, fileName))

	if err != nil {
		panic(fmt.Sprintf("Cannot open file %s", err))
	}

	lines := strings.Split(string(file), "\n")

	return lines
}

func ConvertToCharArray(lines []string) [][]rune {
	res := make([][]rune, len(lines))

	for i, line := range lines {
		res[i] = []rune(line)
		for j, c := range line {
			res[i][j] = c
		}
	}

	return res
}

type Coordinate struct {
	X, Y int
}

type Line struct {
	Start, End Coordinate
}

func (l Line) Value() float64 {
	x := math.Pow(float64(l.End.X-l.Start.X), 2)
	y := math.Pow(float64(l.End.Y-l.Start.Y), 2)
	return math.Sqrt(x+y) + 1
}

func (l Line) Crosses(other Line) bool {
	o1 := orientation(l.Start, l.End, other.Start)
	o2 := orientation(l.Start, l.End, other.End)
	o3 := orientation(other.Start, other.End, l.Start)
	o4 := orientation(other.Start, other.End, l.End)

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == 0 && onSegment(l.Start, other.Start, l.End) {
		return true
	}

	if o2 == 0 && onSegment(l.Start, other.End, l.End) {
		return true
	}

	if o3 == 0 && onSegment(other.Start, l.Start, other.End) {
		return true
	}

	if o4 == 0 && onSegment(other.Start, l.End, other.End) {
		return true
	}

	return false
}

func onSegment(p Coordinate, q Coordinate, r Coordinate) bool {
	if float64(q.X) <= math.Max(float64(p.X), float64(r.X)) &&
		float64(q.X) >= math.Min(float64(p.X), float64(r.X)) &&
		float64(q.Y) <= math.Max(float64(p.Y), float64(r.Y)) &&
		float64(q.Y) >= math.Min(float64(p.Y), float64(r.Y)) {
		return true
	}

	return false
}

func orientation(p Coordinate, q Coordinate, r Coordinate) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	if val == 0 {
		return 0
	}

	if val > 0 {
		return 1
	} else {
		return 2
	}
}

const (
	Top = iota
	Bot
	Left
	Right
	TopLeft
	TopRight
	BotLeft
	BotRight
)

func getNewCoordinatesFromDirection(i int, j int, dir int) (int, int) {
	switch dir {
	case Top:
		return i - 1, j
	case Bot:
		return i + 1, j
	case Left:
		return i, j - 1
	case Right:
		return i, j + 1
	case TopLeft:
		return i - 1, j - 1
	case TopRight:
		return i - 1, j + 1
	case BotLeft:
		return i + 1, j - 1
	case BotRight:
		return i + 1, j + 1
	}

	panic(fmt.Sprintf("Received direction %d", dir))
}

func validateBounds(i int, j int, height int, width int) bool {
	return i >= 0 && i < width && j >= 0 && j < height
}
