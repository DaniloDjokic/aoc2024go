package day13

import (
	"aoc2024go/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day13 struct{}

func (d Day13) Id() int {
	return 13
}

func (d Day13) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

type Machine struct {
	BtnAOffsetX, BtnAOffsetY, BtnBOffsetX, BtnBOffsetY, PrizeX, PrizeY int
}

func (d Day13) Part1(lines []string) (int, error) {
	machines := parseMachines(lines)

	total := 0

	for _, m := range machines {
		xEquation := SolveDiophantine(m.BtnAOffsetX, m.BtnBOffsetX, m.PrizeX)

		if xEquation.HasSolution {
			solutions := xEquation.GetSolutionInRange(1, 100)
			for _, s := range solutions {
				fmt.Println(s)
			}
		}
	}

	return total, nil
}

func parseMachines(lines []string) []Machine {
	machines := make([]Machine, 0)

	machine := Machine{}
	rowNum := 1
	for _, line := range lines {
		if len(line) == 0 {
			machines = append(machines, machine)
			machine = Machine{}
			rowNum = 1
			continue
		}

		switch rowNum {
		case 1:
			plus1 := strings.Index(line, "+")
			plus2 := strings.LastIndex(line, "+")
			comma := strings.Index(line, ",")
			offsetXString := line[plus1+1 : comma]
			offsetYString := line[plus2+1:]
			offsetX, _ := strconv.Atoi(offsetXString)
			offsetY, _ := strconv.Atoi(offsetYString)
			machine.BtnAOffsetX = offsetX
			machine.BtnAOffsetY = offsetY
		case 2:
			plus1 := strings.Index(line, "+")
			plus2 := strings.LastIndex(line, "+")
			comma := strings.Index(line, ",")
			offsetXString := line[plus1+1 : comma]
			offsetYString := line[plus2+1:]
			offsetX, _ := strconv.Atoi(offsetXString)
			offsetY, _ := strconv.Atoi(offsetYString)
			machine.BtnBOffsetX = offsetX
			machine.BtnBOffsetY = offsetY
		case 3:
			equals1 := strings.Index(line, "=")
			equals2 := strings.LastIndex(line, "=")
			comma := strings.Index(line, ",")
			prizeXString := line[equals1+1 : comma]
			prizeYString := line[equals2+1:]
			prizeX, _ := strconv.Atoi(prizeXString)
			prizeY, _ := strconv.Atoi(prizeYString)
			machine.PrizeX = prizeX
			machine.PrizeY = prizeY
		}

		rowNum++
	}

	//for _, machine := range machines {
	//	fmt.Println(machine)
	//}

	return machines
}

func extendedGCD(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	d, x1, y1 := extendedGCD(b, a%b)
	return d, y1, x1 - (a/b)*y1
}

type Solution struct {
	X0, Y0       int
	XStep, YStep int
	HasSolution  bool
}

func SolveDiophantine(a, b, c int) Solution {
	gcd, x0, y0 := extendedGCD(a, b)

	if c%gcd != 0 {
		return Solution{HasSolution: false}
	}

	factor := c / gcd
	return Solution{
		X0:          x0 * factor,
		Y0:          y0 * factor,
		XStep:       b / gcd,
		YStep:       -a / gcd,
		HasSolution: true,
	}
}

func (s Solution) GetSolution(t int) (x, y int) {
	if !s.HasSolution {
		return 0, 0
	}

	return s.X0 + s.XStep*t, s.Y0 + s.YStep*t
}

func (s Solution) GetSolutionInRange(minXY, maxXY int) []struct{ X, Y int } {
	if !s.HasSolution {
		return nil
	}

	var solutions []struct{ X, Y int }

	tMin := max((minXY-s.X0)/s.XStep, (s.Y0-maxXY)/s.YStep)
	tMax := min((maxXY-s.X0)/s.XStep, (s.Y0-minXY)/s.YStep)

	for t := tMin; t <= tMax; t++ {
		x, y := s.GetSolution(t)
		if x >= minXY && x <= maxXY && y >= minXY && y <= maxXY {
			solutions = append(solutions, struct{ X, Y int }{x, y})
		}
	}

	return solutions
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (d Day13) Part2(lines []string) (int, error) { panic("A") }
