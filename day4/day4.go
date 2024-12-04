package day4

import (
	"aoc2024go/utils"
	"fmt"
)

type Day4 struct{}

func (d Day4) Id() int {
	return 4
}

func (d Day4) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day4) Part1(lines []string) (int, error) {
	matrix := utils.ConvertToCharArray(lines)
	total := 0
	height := len(matrix)
	width := len(matrix[0])

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 'X' {
				total += startSearch(i, j, matrix, height, width)
			}
		}
	}

	return total, nil
}

func startSearch(i int, j int, matrix [][]rune, height int, width int) int {
	total := 0

	total += search(i, j, height, width, matrix, Left)
	total += search(i, j, height, width, matrix, Right)
	total += search(i, j, height, width, matrix, Top)
	total += search(i, j, height, width, matrix, Bot)
	total += search(i, j, height, width, matrix, TopLeft)
	total += search(i, j, height, width, matrix, TopRight)
	total += search(i, j, height, width, matrix, BotLeft)
	total += search(i, j, height, width, matrix, BotRight)

	return total
}

func search(i int, j int, height int, width int, matrix [][]rune, dir int) int {
	xmas := "XMAS"
	total := 0
	count := 1
	word := make([]rune, 4)
	word[0] = 'X'

	newI, newJ := getNewCoordinatesFromDirection(i, j, dir)

	for validateBounds(newI, newJ, height, width) {
		nextLetter := matrix[newI][newJ]
		word[count] = nextLetter

		if count == 3 {
			if string(word) == xmas {
				total++
			}
			break
		}

		newI, newJ = getNewCoordinatesFromDirection(newI, newJ, dir)
		count++
	}

	return total
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

func (d Day4) Part2(lines []string) (int, error) { panic("Not implemented") }
