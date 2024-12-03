package day3

import "aoc2024go/utils"

type Day3 struct{}

func (d Day3) Id() int {
	return 3
}

func (d Day3) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day3) Part1(lines []string) (int, error) {
	p := NewParser()

	for _, line := range lines {
		p.Parse(line)
	}

	vals := p.results
	acc := 0

	for _, val := range vals {
		acc += val.val
	}

	return acc, nil
}

func (d Day3) Part2(lines []string) (int, error) {
	p := NewParser()

	for _, line := range lines {
		p.Parse(line)
	}

	vals := p.results
	acc := 0

	for _, val := range vals {
		if val.enabled {
			acc += val.val
		}
	}

	return acc, nil
}
