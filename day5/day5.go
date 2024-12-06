package day5

import (
	"aoc2024go/utils"
	"strconv"
	"strings"
)

type Day5 struct{}

func (d Day5) Id() int {
	return 5
}

func (d Day5) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day5) Part1(lines []string) (int, error) {
	comp, br := d.ProcessMapping(lines)
	total := 0

	for i := br + 1; i < len(lines); i++ { //Process commands
		line := lines[i]
		split := strings.Split(line, ",")
		validLine := true

		alreadySeen := make(map[int]bool)

		for i := 0; i < len(split); i++ { // loop comma delimiter numbers
			numS := split[i]
			left, _ := strconv.Atoi(numS)

			biggerThanSet, ok := comp[left]
			if ok {
				for key, _ := range biggerThanSet {
					if _, ok := alreadySeen[key]; ok {
						validLine = false
					}
				}
			}
			alreadySeen[left] = true
		}

		if validLine {
			middle := split[len(split)/2]
			middleNum, _ := strconv.Atoi(middle)
			total += middleNum
		}
	}

	return total, nil
}

func (d Day5) ProcessMapping(lines []string) (map[int]map[int]bool, int) {
	comp := make(map[int]map[int]bool)

	for i, line := range lines {
		if line != "" {
			split := strings.Split(line, "|")
			left, _ := strconv.Atoi(strings.TrimSpace(split[0]))
			right, _ := strconv.Atoi(strings.TrimSpace(split[1]))

			val, ok := comp[left]
			if !ok {
				val = make(map[int]bool)
				val[right] = true
				comp[left] = val
			} else {
				val[right] = true
				comp[left] = val
			}
		} else {
			return comp, i
		}
	}

	panic("How?")
}

func (d Day5) Part2(lines []string) (int, error) { panic("Not implemented") }
