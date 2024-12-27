package day12

import (
	"aoc2024go/utils"
	"fmt"
)

type Day12 struct{}

func (d Day12) Id() int {
	return 12
}

func (d Day12) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

type Crop struct {
	Symbol    rune
	X, Y      int
	Processed bool
}

func printGroups(groups [][]*Crop, dimX, dimY int, grid [][]Crop) {
	for _, group := range groups {
		fmt.Println(fmt.Sprintf("Group: %s", string(group[0].Symbol)))
		for _, crop := range group {
			fmt.Println(*crop)
		}

		fmt.Println(fmt.Sprintf("Area: %d", calcGroupArea(group)))
		fmt.Println(fmt.Sprintf("Parimeter: %d", calcGroupPerimeter(group, dimX, dimY, grid)))
	}
}

func (d Day12) Part1(lines []string) (int, error) {
	charArray := utils.ConvertToCharArray(lines)

	grid := make([][]Crop, len(charArray))

	for i := range charArray {
		grid[i] = make([]Crop, len(charArray[i]))
		for j, c := range charArray[i] {
			grid[i][j] = Crop{Symbol: c, X: i, Y: j, Processed: false}
		}
	}

	dimX := len(grid)
	dimY := len(grid[0])

	groups := make([][]*Crop, 0)

	for i := range grid {
		for j := range grid[i] {
			if !grid[i][j].Processed {
				group := createGroup(&grid[i][j], i, j, dimX, dimY, grid)
				groups = append(groups, group)
			}
		}
	}

	//printGroups(groups, dimX, dimY, grid)

	total := 0
	for _, group := range groups {
		area := calcGroupArea(group)
		parimeter := calcGroupPerimeter(group, dimX, dimY, grid)
		total += area * parimeter
	}

	return total, nil
}

func calcGroupArea(group []*Crop) int {
	return len(group)
}

func calcGroupPerimeter(group []*Crop, dimX, dimY int, grid [][]Crop) int {
	total := 0
	for _, crop := range group {
		initial := 4

		i := crop.X
		j := crop.Y

		if isInGroup(i+1, j, dimX, dimY, group, grid) {
			initial--
		}
		if isInGroup(i-1, j, dimX, dimY, group, grid) {
			initial--
		}
		if isInGroup(i, j+1, dimX, dimY, group, grid) {
			initial--
		}
		if isInGroup(i, j-1, dimX, dimY, group, grid) {
			initial--
		}

		total += initial
	}

	return total
}

func isInGroup(i, j, dimX, dimY int, group []*Crop, grid [][]Crop) bool {
	if utils.IsOffMapRaw(i, j, dimX, dimY) {
		return false
	}

	if groupContains(group, &grid[i][j]) {
		return true
	}

	return false
}

func groupContains(group []*Crop, crop *Crop) bool {
	for _, c := range group {
		if c == crop {
			return true
		}
	}

	return false
}

func createGroup(start *Crop, i, j, dimX, dimY int, grid [][]Crop) []*Crop {
	group := make([]*Crop, 0)

	group = append(group, start)
	start.Processed = true

	traverseGrid(start, i+1, j, dimX, dimY, grid, &group)
	traverseGrid(start, i-1, j, dimX, dimY, grid, &group)
	traverseGrid(start, i, j+1, dimX, dimY, grid, &group)
	traverseGrid(start, i, j-1, dimX, dimY, grid, &group)

	return group
}

func traverseGrid(prev *Crop, i, j, dimX, dimY int, grid [][]Crop, group *[]*Crop) {
	if utils.IsOffMapRaw(i, j, dimX, dimY) {
		return
	}

	currCrop := &grid[i][j]
	if prev.Symbol == currCrop.Symbol && !currCrop.Processed {
		*group = append(*group, currCrop)

		currCrop.Processed = true

		traverseGrid(currCrop, i+1, j, dimX, dimY, grid, group)
		traverseGrid(currCrop, i-1, j, dimX, dimY, grid, group)
		traverseGrid(currCrop, i, j+1, dimX, dimY, grid, group)
		traverseGrid(currCrop, i, j-1, dimX, dimY, grid, group)
	}
}

func (d Day12) Part2(lines []string) (int, error) { panic("AAAAA") }
