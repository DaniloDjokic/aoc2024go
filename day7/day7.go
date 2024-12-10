package day7

import (
	"aoc2024go/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct{}

func (d Day7) Id() int {
	return 7
}

func (d Day7) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day7) Part1(lines []string) (int, error) {
	var total int64
	total = 0

	for _, line := range lines {
		split := strings.Split(line, ":")
		expectedRes, _ := strconv.ParseInt(split[0], 10, 64)

		//fmt.Println(fmt.Sprintf("Expected: %d", expectedRes))

		equationInputs := strings.Split(split[1], " ")
		firstVal, _ := strconv.ParseInt(equationInputs[1], 10, 64)

		inputs := make([]int64, 0)

		for _, elStr := range equationInputs[2:] {
			el, _ := strconv.ParseInt(elStr, 10, 64)
			inputs = append(inputs, el)
		}

		if buildTree(firstVal, inputs, 0, expectedRes) {
			total += expectedRes
		}
	}

	fmt.Println(fmt.Sprintf("Actual result that didn't overflow: %d", total))
	return 0, nil
}

func buildTree(curr int64, vals []int64, level int, target int64) bool {
	if level == len(vals) {
		if curr == target {
			return true
		} else {
			return false
		}
	}

	if curr > target {
		return false
	}

	next := vals[level]

	leftVal := calcRes(curr, next, Add)
	midVal := calcRes(curr, next, Concat)
	rightVal := calcRes(curr, next, Mul)

	left := buildTree(leftVal, vals, level+1, target)
	mid := buildTree(midVal, vals, level+1, target)
	right := buildTree(rightVal, vals, level+1, target)

	return left || mid || right
}

func calcRes(a, b, op int64) int64 {
	switch op {
	case Add:
		return a + b
	case Mul:
		return a * b
	case Concat:
		s := strconv.FormatInt(a, 10) + strconv.FormatInt(b, 10)
		n, _ := strconv.ParseInt(s, 10, 64)

		return n
	}

	panic("Receiver invalid operation")
}

const (
	Add = iota
	Mul
	Concat
)

func (d Day7) Part2(lines []string) (int, error) { panic("Not implemented") }
