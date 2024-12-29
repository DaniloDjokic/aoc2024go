package day12

import (
	"aoc2024go/utils"
	"fmt"
	"sort"
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

	printGroups(groups, dimX, dimY, grid)

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
	edges := make(map[int][]utils.Coordinate)

	for _, crop := range group {
		i := crop.X
		j := crop.Y

		//Bot
		if !isInGroup(i+1, j, dimX, dimY, group, grid) {
			if dirEdges, ok := edges[utils.Bot]; ok {
				edges[utils.Bot] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			} else {
				edges[utils.Bot] = make([]utils.Coordinate, 0)
				edges[utils.Bot] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			}
		}
		//Top
		if !isInGroup(i-1, j, dimX, dimY, group, grid) {
			if dirEdges, ok := edges[utils.Top]; ok {
				edges[utils.Top] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			} else {
				edges[utils.Top] = make([]utils.Coordinate, 0)
				edges[utils.Top] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			}
		}
		//Right
		if !isInGroup(i, j+1, dimX, dimY, group, grid) {
			if dirEdges, ok := edges[utils.Right]; ok {
				edges[utils.Right] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			} else {
				edges[utils.Right] = make([]utils.Coordinate, 0)
				edges[utils.Right] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			}
		}
		//Left
		if !isInGroup(i, j-1, dimX, dimY, group, grid) {
			if dirEdges, ok := edges[utils.Left]; ok {
				edges[utils.Left] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			} else {
				edges[utils.Left] = make([]utils.Coordinate, 0)
				edges[utils.Left] = append(dirEdges, utils.Coordinate{X: i, Y: j})
			}
		}
	}

	for dir, edgesByDir := range edges {
		dirEdgesByAxis := make(map[int]*[]int)

		if dir == utils.Top || dir == utils.Bot {
			for _, edge := range edgesByDir {
				if row, ok := dirEdgesByAxis[edge.X]; !ok {
					newRow := make([]int, 0)
					newRow = append(newRow, edge.Y)
					dirEdgesByAxis[edge.X] = &newRow
				} else {
					*row = append(*row, edge.Y)
				}
			}
		} else {
			for _, edge := range edgesByDir {
				if row, ok := dirEdgesByAxis[edge.Y]; !ok {
					newRow := make([]int, 0)
					newRow = append(newRow, edge.X)
					dirEdgesByAxis[edge.Y] = &newRow
				} else {
					*row = append(*row, edge.X)
				}
			}
		}

		for _, axis := range dirEdgesByAxis {
			sort.Ints(*axis)

			lineCount := 1

			for i := 1; i < len(*axis); i++ {
				if (*axis)[i]-(*axis)[i-1] > 1 {
					lineCount++
				}
			}

			total += lineCount
		}
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
