package utils

import (
	"fmt"
	"os"
	"strings"
)

type Day interface {
	Id() int
	LoadFile(isTest bool) []string
	Run1(lines []string) (int, error)
	Run2(lines []string) (int, error)
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
