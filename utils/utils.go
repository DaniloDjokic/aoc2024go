package utils

import (
	"fmt"
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
