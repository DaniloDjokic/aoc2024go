package day8

import (
	"aoc2024go/utils"
	"fmt"
	"math"
)

type Day8 struct{}

func (d Day8) Id() int {
	return 8
}

func (d Day8) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day8) Part1(lines []string) (int, error) {
	matrix := utils.ConvertToCharArray(lines)
	antennas := make(map[rune][]utils.Coordinate)

	antinodes := make(map[utils.Coordinate]bool)
	total := 0

	dimX := len(matrix)
	dimY := len(matrix[0])

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != '.' {
				antennaChar := matrix[i][j]

				if _, ok := antennas[antennaChar]; !ok {
					antennas[antennaChar] = make([]utils.Coordinate, 0)
				}

				antennas[antennaChar] = append(antennas[antennaChar], utils.Coordinate{X: i, Y: j})
			}
		}
	}

	for _, antennaCoords := range antennas {
		for _, c1 := range antennaCoords {
			for _, c2 := range antennaCoords {
				if c1.X == c2.X && c1.Y == c2.Y {
					continue
				}

				offsetX := float64(c2.X - c1.X)
				offsetY := float64(c2.Y - c1.Y)

				nextPoint, offMap := calcNextPoint(c2, int(offsetX), int(offsetY), dimX, dimY)

				if !offMap {
					if _, ok := antinodes[nextPoint]; !ok {
						antinodes[nextPoint] = true
						total++
					}
				}
			}
		}
	}

	for k, _ := range antinodes {
		fmt.Println(k)
	}

	return total, nil
}

func calcNextPoint(start utils.Coordinate, offsetX, offsetY, dimX, dimY int) (utils.Coordinate, bool) {
	nextPoint := utils.Coordinate{X: start.X + offsetX, Y: start.Y + offsetY}

	if utils.IsOffMap(nextPoint, dimX, dimY) {
		return utils.Coordinate{}, true
	}

	return nextPoint, false
}

func (d Day8) Part2(lines []string) (int, error) {
	matrix := utils.ConvertToCharArray(lines)
	antennas := make(map[rune][]utils.Coordinate)

	antinodes := make(map[utils.Coordinate]bool)
	total := 0

	dimX := len(matrix)
	dimY := len(matrix[0])

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != '.' {
				antennaChar := matrix[i][j]

				if _, ok := antennas[antennaChar]; !ok {
					antennas[antennaChar] = make([]utils.Coordinate, 0)
				}

				antennas[antennaChar] = append(antennas[antennaChar], utils.Coordinate{X: i, Y: j})
			}
		}
	}

	for _, antennaCoords := range antennas {
		for _, c1 := range antennaCoords {
			for _, c2 := range antennaCoords {
				if c1.X == c2.X && c1.Y == c2.Y {
					continue
				}

				nextPoint := utils.Coordinate{X: c2.X, Y: c2.Y}
				if _, ok := antinodes[nextPoint]; !ok {
					antinodes[nextPoint] = true
					total++
				}

				offsetX := float64(c2.X - c1.X)
				offsetY := float64(c2.Y - c1.Y)

				if math.Abs(offsetX) == math.Abs(offsetY) {
					offsetX /= math.Abs(offsetX)
					offsetY /= math.Abs(offsetY)
				}

				var offMap bool
				nextPoint, offMap = calcNextPoint(c2, int(offsetX), int(offsetY), dimX, dimY)

				for !offMap {
					if _, ok := antinodes[nextPoint]; !ok {
						antinodes[nextPoint] = true
						total++
					}

					nextPoint, offMap = calcNextPoint(nextPoint, int(offsetX), int(offsetY), dimX, dimY)
				}
			}
		}
	}

	for k, _ := range antinodes {
		fmt.Println(k)
	}

	return total, nil
}
