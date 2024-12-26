package day11

import (
	"aoc2024go/utils"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Day11 struct{}

func (d Day11) Id() int {
	return 11
}

func (d Day11) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day11) Part1(lines []string) (int, error) {
	stones := make([]int, 0)

	for _, c := range strings.Split(lines[0], " ") {
		num, _ := strconv.Atoi(string(c))
		stones = append(stones, num)
	}

	wg := sync.WaitGroup{}
	total := 0

	for _, s := range stones {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newTotal := processStone(Stone{val: s, depth: 0})
			total += newTotal
		}()
	}

	wg.Wait()

	return total, nil
}

func processStone(s Stone) int {
	st := stack.New()
	st.Push(s)
	total := 1

	stoneCache := make(map[int]StoneCacheVal)

	for i := 1; i <= 6; i++ {
		nextNums := make([]Stone, 0, st.Len())

		for st.Len() != 0 {
			nextNums = append(nextNums, st.Pop().(Stone))
		}

		for _, curr := range nextNums {

			if curr.val == 0 {
				st.Push(Stone{val: 1, depth: curr.depth + 1})
				continue
			}

			digitCount := getDigitCount(curr.val)
			if digitCount%2 == 0 {
				divisor := int(math.Pow(10, math.Floor(float64(digitCount)/2)))

				left := removeLeadingZerosMath(int(math.Floor(float64(curr.val) / float64(divisor))))
				right := removeLeadingZerosMath(curr.val % divisor)

				st.Push(Stone{val: left, depth: curr.depth + 1})
				st.Push(Stone{val: right, depth: curr.depth + 1})

				total++
			} else {
				st.Push(Stone{val: curr.val * 2024, depth: curr.depth + 1})
			}
		}
	}

	fmt.Println(stoneCache)

	return total
}

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

type StoneCacheVal struct {
	DepthAtStart, DepthToAdd, SplitsAtStart, SplitsToAdd int
	Processed                                            bool
	SubTreeLeaves                                        []SubTreeLeaf
}

func (d Day11) Part2(lines []string) (int, error) {
	stones := make([]int, 0)

	for _, c := range strings.Split(lines[0], " ") {
		num, _ := strconv.Atoi(c)
		stones = append(stones, num)
	}

	total := 0

	var cache map[int]StoneCacheVal
	for _, s := range stones {
		total++
		cache = make(map[int]StoneCacheVal)
		allBonus := make([]SubTreeLeaf, 0)
		processStoneRecursive(Stone{val: s, depth: 0}, 0, 25, cache, false, &total, &allBonus)

		temp := make([]SubTreeLeaf, 0)

		for _, curr := range allBonus {
			processStoneRecursive(curr.Stone, curr.Splits, 25, cache, false, &total, &temp)
		}
	}

	return total, nil
}

type SubTreeLeaf struct {
	Stone  Stone
	Splits int
}

func createSubtree(s Stone, startIdx, maxDepth, splits int, bonus *[]SubTreeLeaf) {
	if s.depth == maxDepth {
		//if s.val != startIdx {
		*bonus = append(*bonus, SubTreeLeaf{Stone: Stone{val: s.val, depth: s.depth}, Splits: splits})
		//}
		return
	}

	if s.val == 0 {
		createSubtree(Stone{val: 1, depth: s.depth + 1}, startIdx, maxDepth, splits, bonus)
		return
	}

	digitCount := getDigitCount(s.val)
	if digitCount%2 == 0 {
		divisor := int(math.Pow(10, math.Floor(float64(digitCount)/2)))

		left := removeLeadingZerosMath(int(math.Floor(float64(s.val) / float64(divisor))))
		leftStone := Stone{val: left, depth: s.depth + 1}
		right := removeLeadingZerosMath(s.val % divisor)
		rightStone := Stone{val: right, depth: s.depth + 1}

		createSubtree(leftStone, startIdx, maxDepth, splits+1, bonus)
		createSubtree(rightStone, startIdx, maxDepth, splits+1, bonus)
	} else {
		createSubtree(Stone{val: s.val * 2024, depth: s.depth + 1}, startIdx, maxDepth, splits, bonus)
	}
}

func processStoneRecursive(s Stone, splits int, maxDepth int, cache map[int]StoneCacheVal, ignoreCache bool, total *int, allBonus *[]SubTreeLeaf) {
	if s.depth == maxDepth {
		return
	}

	if !ignoreCache {
		stone, ok := cache[s.val]
		if !ok {
			cache[s.val] = StoneCacheVal{DepthAtStart: s.depth, DepthToAdd: 0, SplitsAtStart: splits, SplitsToAdd: 0}
		} else {
			if s.depth > stone.DepthAtStart {
				if !stone.Processed {
					stone.DepthToAdd = s.depth - stone.DepthAtStart
					stone.SplitsToAdd = splits - stone.SplitsAtStart

					subTreeLeaves := make([]SubTreeLeaf, 0)
					createSubtree(Stone{val: s.val, depth: stone.DepthAtStart}, s.val, s.depth, stone.SplitsAtStart, &subTreeLeaves)
					stone.SubTreeLeaves = subTreeLeaves

					stone.Processed = true
					cache[s.val] = stone
				}

				bonuses := processMemo(s, stone, splits, maxDepth, cache, total)

				for _, bs := range bonuses {
					*allBonus = append(*allBonus, bs)
				}

				return
			}
		}
	}

	if s.val == 0 {
		processStoneRecursive(Stone{val: 1, depth: s.depth + 1}, splits, maxDepth, cache, ignoreCache, total, allBonus)
		return
	}

	digitCount := getDigitCount(s.val)
	if digitCount%2 == 0 {
		divisor := int(math.Pow(10, math.Floor(float64(digitCount)/2)))

		left := removeLeadingZerosMath(int(math.Floor(float64(s.val) / float64(divisor))))
		leftStone := Stone{val: left, depth: s.depth + 1}
		right := removeLeadingZerosMath(s.val % divisor)
		rightStone := Stone{val: right, depth: s.depth + 1}

		*total++

		processStoneRecursive(leftStone, splits+1, maxDepth, cache, ignoreCache, total, allBonus)
		processStoneRecursive(rightStone, splits+1, maxDepth, cache, ignoreCache, total, allBonus)
	} else {
		processStoneRecursive(Stone{val: s.val * 2024, depth: s.depth + 1}, splits, maxDepth, cache, ignoreCache, total, allBonus)
	}
}

func processMemo(s Stone, stone StoneCacheVal, splits, maxDepth int, cache map[int]StoneCacheVal, total *int) []SubTreeLeaf {
	remainingDepth := maxDepth - s.depth
	repeatCount := remainingDepth / stone.DepthToAdd

	allBonus := make([]SubTreeLeaf, 0)
	if repeatCount < 1 {
		emptyBonus := make([]SubTreeLeaf, 0)
		processStoneRecursive(s, splits, maxDepth, cache, true, total, &emptyBonus)
	} else {
		*total += repeatCount * (len(stone.SubTreeLeaves)) // +1 To offset that we don't add the same index to subtree

		for _, b := range stone.SubTreeLeaves {
			for i := 1; i <= repeatCount; i++ {
				bonusStone := SubTreeLeaf{Stone: Stone{val: b.Stone.val, depth: s.depth + (stone.DepthToAdd * i)}, Splits: 1}
				allBonus = append(allBonus, bonusStone)
			}
		}
	}

	return allBonus
}
