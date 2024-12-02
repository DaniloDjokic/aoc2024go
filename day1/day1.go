package day1

import (
	"aoc2024go/utils"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d Day1) Id() int {
	return 1
}

func (d Day1) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, 1)
}

func (d Day1) Part1(lines []string) (int, error) {
	panic("Not implemented")
}

func (d Day1) Part2(lines []string) (int, error) {
	var leftInts []int
	rightInts := make(map[int]int)

	for _, line := range lines {

		nums := strings.Split(line, " ")

		left, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		right, errr := strconv.Atoi(nums[3])
		if errr != nil {
			panic(errr)
		}

		leftInts = append(leftInts, left)
		_, ok := rightInts[right]

		if ok {
			rightInts[right]++
		} else {
			rightInts[right] = 1
		}
	}

	total := 0

	for _, left := range leftInts {
		simCore, ok := rightInts[left]

		var newVal int

		if ok {
			newVal = left * simCore
		} else {
			newVal = 0
		}

		total += newVal
	}

	return total, nil
}
