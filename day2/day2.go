package day2

import (
	"aoc2024go/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day2 struct{}

func (d Day2) Id() int {
	return 2
}

func (d Day2) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, 2)
}

func (d Day2) Part1(lines []string) (int, error) {
	totalSafety := 0
	for _, line := range lines {
		numStrings := strings.Split(line, " ")
		nums := make([]int, len(numStrings))

		var err error
		for i, s := range numStrings {
			nums[i], err = strconv.Atoi(s)

			if err != nil {
				panic(fmt.Sprintf("Somehow, %s is not an int", s))
			}
		}

		safety := getSafety(nums)
		totalSafety += safety
	}

	return totalSafety, nil
}

func getSafety(nums []int) int {
	dir := 0 //-1, 0, 1
	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		curr := nums[i]

		if prev == curr {
			return 0
		}

		diff := int(math.Abs(float64(prev - curr)))
		if diff < 1 || diff > 3 {
			return 0
		}

		if prev < curr {
			if dir == -1 {
				return 0
			}
			dir = 1
		} else if prev > curr {
			if dir == 1 {
				return 0
			}
			dir = -1
		}

		prev = curr
	}

	return 1
}

func (d Day2) Part2(lines []string) (int, error) {
	panic("Not implemented")
}
