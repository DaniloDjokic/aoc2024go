package day13

import (
	"aoc2024go/utils"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
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

	for i, m := range machines {
		a, b, err := m.solve()
		if err != nil {
			fmt.Printf("Machine %d has no solution, %s\n", i+1, err.Error())
		} else {
			fmt.Printf("Machine %d: A=%d, B=%d\n", i+1, a, b)
			total += 3*a + b
		}
	}

	return total, nil
}

func parseMachines(lines []string) []Machine {
	machines := make([]Machine, 0)

	machine := Machine{}
	rowNum := 1

	for i, line := range lines {
		if i == len(lines)-1 {
			continue
		}

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

	return machines
}

func (m *Machine) solve() (int, int, error) {
	A := mat.NewDense(
		2,
		2,
		[]float64{
			float64(m.BtnAOffsetX),
			float64(m.BtnBOffsetX),
			float64(m.BtnAOffsetY),
			float64(m.BtnBOffsetY),
		},
	)
	B := mat.NewVecDense(2, []float64{float64(m.PrizeX), float64(m.PrizeY)})

	var x mat.VecDense
	if err := x.SolveVec(A, B); err != nil {
		panic(err)
	}

	for _, v := range x.RawVector().Data {
		if v < 0 {
			return 0, 0, errors.New("No solution for equation, number is negative")
		}

		if !isWholeNumber(v, 1e-6) {
			return 0, 0, errors.New("No solution for equation")
		}

		if v > 100 {
			return 0, 0, errors.New("No solution smaller than 101")
		}
	}

	return int(math.Round(x.At(0, 0))), int(math.Round(x.At(1, 0))), nil
}

func isWholeNumber(f float64, tolerance float64) bool {
	// rounded := math.Round(f)
	// left := f - rounded
	// res := math.Abs(left) < tolerance
	// return res
	return math.Abs(f-math.Round(f)) < tolerance
}

func (d Day13) Part2(lines []string) (int, error) { panic("A") }
