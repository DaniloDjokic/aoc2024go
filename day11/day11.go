package day11

import (
	"aoc2024go/utils"
	"math"
	"strconv"
	"strings"
)

type Day11 struct{}

func (d Day11) Id() int {
	return 11
}

func (d Day11) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day11) Part1(lines []string) (int, error) {
	//stones := make([]int, 0)
	//
	//for _, c := range strings.Split(lines[0], " ") {
	//	num, _ := strconv.Atoi(string(c))
	//	stones = append(stones, num)
	//}
	//
	//wg := sync.WaitGroup{}
	//total := 0
	//
	//for _, s := range stones {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		newTotal := processStone(Stone{val: s, depth: 0})
	//		total += newTotal
	//	}()
	//}
	//
	//wg.Wait()

	return 0, nil
}

//func processStone(s Stone) int {
//	st := stack.New()
//	st.Push(s)
//	total := 1
//
//	stoneCache := make(map[int]StoneCacheVal)
//
//	for i := 1; i <= 6; i++ {
//		nextNums := make([]Stone, 0, st.Len())
//
//		for st.Len() != 0 {
//			nextNums = append(nextNums, st.Pop().(Stone))
//		}
//
//		for _, curr := range nextNums {
//
//			if curr.val == 0 {
//				st.Push(Stone{val: 1, depth: curr.depth + 1})
//				continue
//			}
//
//			digitCount := getDigitCount(curr.val)
//			if digitCount%2 == 0 {
//				divisor := int(math.Pow(10, math.Floor(float64(digitCount)/2)))
//
//				left := removeLeadingZerosMath(int(math.Floor(float64(curr.val) / float64(divisor))))
//				right := removeLeadingZerosMath(curr.val % divisor)
//
//				st.Push(Stone{val: left, depth: curr.depth + 1})
//				st.Push(Stone{val: right, depth: curr.depth + 1})
//
//				total++
//			} else {
//				st.Push(Stone{val: curr.val * 2024, depth: curr.depth + 1})
//			}
//		}
//	}
//
//	fmt.Println(stoneCache)
//
//	return total
//}

func getDigitCount(n int) int {
	return int(math.Floor(math.Log10(float64(n))) + 1)
}
func removeLeadingZerosMath(number int) int {
	if number == 0 {
		return 0
	}

	digits := int(math.Log10(float64(number))) + 1

	return number % int(math.Pow10(digits))
}

type Stone struct {
	val, depth int
}

func (d Day11) Part2(lines []string) (int, error) {
	stones := make([]int, 0)

	for _, c := range strings.Split(lines[0], " ") {
		num, _ := strconv.Atoi(c)
		stones = append(stones, num)
	}

	total := 0
	for _, s := range stones {
		startStone := Stone{val: s, depth: 0}
		memo := make(map[Stone]int)
		total += processStoneRecursive(startStone, 75, memo)
	}

	return total, nil
}

func processStoneRecursive(s Stone, maxDepth int, memo map[Stone]int) int {
	if res, ok := memo[s]; ok {
		return res
	}

	if s.depth >= maxDepth {
		return 1
	}

	if s.val == 0 {
		memo[s] = processStoneRecursive(Stone{val: 1, depth: s.depth + 1}, maxDepth, memo)
		return memo[s]
	}

	digitCount := getDigitCount(s.val)
	if digitCount%2 == 0 {
		divisor := int(math.Pow(10, math.Floor(float64(digitCount)/2)))

		left := removeLeadingZerosMath(int(math.Floor(float64(s.val) / float64(divisor))))
		leftStone := Stone{val: left, depth: s.depth + 1}
		right := removeLeadingZerosMath(s.val % divisor)
		rightStone := Stone{val: right, depth: s.depth + 1}

		memo[s] = processStoneRecursive(leftStone, maxDepth, memo) + processStoneRecursive(rightStone, maxDepth, memo)
		return memo[s]
	} else {
		memo[s] = processStoneRecursive(Stone{val: s.val * 2024, depth: s.depth + 1}, maxDepth, memo)
		return memo[s]
	}
}
