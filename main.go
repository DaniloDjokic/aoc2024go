package main

import (
	"aoc2024go/day9"
	"aoc2024go/utils"
	"fmt"
	"time"
)

func main() {
	day := day9.Day9{}

	isPart1 := false
	isTest := false

	result := runSolution(day, isTest, isPart1)

	var partLabel, testLabel string

	if isPart1 {
		partLabel = "part 1"
	} else {
		partLabel = "part 2"
	}

	if isTest {
		testLabel = "test input"
	} else {
		testLabel = "real input"
	}

	fmt.Println(fmt.Sprintf("The result for day%d, %s, %s is %d ", day.Id(), partLabel, testLabel, result))
}

func runSolution(day utils.Day, isTest bool, isPart1 bool) int {
	var result int

	start := time.Now()
	lines := day.LoadFile(isTest)
	startFun := time.Now()

	var err error
	if isPart1 {
		result, err = day.Part1(lines)
	} else {
		result, err = day.Part2(lines)
	}

	if err != nil {
		panic(err)
	}

	timeFun := time.Since(startFun)
	timeTotal := time.Since(start)

	fmt.Println("Total exec time is ", timeTotal)
	fmt.Println("Function exec time is", timeFun)

	return result
}
