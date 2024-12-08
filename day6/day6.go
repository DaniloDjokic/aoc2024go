package day6

import (
	"aoc2024go/utils"
)

type Day6 struct{}

func (d Day6) Id() int {
	return 6
}

func (d Day6) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day6) Part1(lines []string) (int, error) {
	matrix := utils.ConvertToCharArray(lines)

	obstacles := make([]utils.Coordinate, 0, 100)
	var start utils.Coordinate
	var offMap bool

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '#' {
				obstacles = append(obstacles, utils.Coordinate{X: i, Y: j})
			}

			if matrix[i][j] == '^' {
				start = utils.Coordinate{X: i, Y: j}
			}
		}
	}

	passedLines := make(map[utils.Line]int)

	nextDir := utils.Top
	var nextStart utils.Coordinate
	_, offMap, nextStart = getNextObstacle(start, obstacles, nextDir)

	passedLine := utils.Line{Start: start, End: nextStart}
	passedLines[passedLine] = int(passedLine.Value())

	for !offMap {
		nextDir = cycleDir(nextDir)
		start = nextStart
		_, offMap, nextStart = getNextObstacle(start, obstacles, nextDir)
		passedLine := utils.Line{Start: start, End: nextStart}
		passedLines[passedLine] = int(passedLine.Value())
	}

	total := 0
	calculatedLines := make([]utils.Line, 0, 10)

	for currLine, val := range passedLines {
		currLineVal := val
		for _, calculatedLine := range calculatedLines {
			if currLine.Crosses(calculatedLine) {
				currLineVal--
			}
		}
		total += currLineVal
		calculatedLines = append(calculatedLines, currLine)
	}

	return total, nil
}

func cycleDir(dir int) int {
	if dir == utils.Top {
		return utils.Right
	}

	if dir == utils.Right {
		return utils.Bot
	}

	if dir == utils.Bot {
		return utils.Left
	}

	if dir == utils.Left {
		return utils.Top
	}

	panic("Received non standard direction cycle")
}

func getNextObstacle(start utils.Coordinate, obstacles []utils.Coordinate, dir int) (utils.Coordinate, bool, utils.Coordinate) {
	var obs, newStart utils.Coordinate
	var found bool

	switch dir {
	case utils.Top:
		obs, found = calcClosestObstacle(start, obstacles, false, false)
		newStart = utils.Coordinate{X: obs.X + 1, Y: obs.Y}
	case utils.Left:
		obs, found = calcClosestObstacle(start, obstacles, true, false)
		newStart = utils.Coordinate{X: obs.X, Y: obs.Y + 1}
	case utils.Right:
		obs, found = calcClosestObstacle(start, obstacles, true, true)
		newStart = utils.Coordinate{X: obs.X, Y: obs.Y - 1}
	case utils.Bot:
		obs, found = calcClosestObstacle(start, obstacles, false, true)
		newStart = utils.Coordinate{X: obs.X - 1, Y: obs.Y}
	default:
		panic("Got direction outside of top bot left right")
	}

	if !found {
		return obs, true, newStart
	} else {
		return obs, false, newStart
	}
}

func calcClosestObstacle(start utils.Coordinate, obstacles []utils.Coordinate, isX bool, obstacleHasBiggerIndex bool) (obstacle utils.Coordinate, found bool) {
	var predicate func(utils.Coordinate, int) bool
	var filterCmp, searchCmp int

	if isX {
		searchCmp = start.Y
		filterCmp = start.X
		predicate = func(c utils.Coordinate, t int) bool {
			return c.X == t
		}
	} else {
		searchCmp = start.X
		filterCmp = start.Y
		predicate = func(c utils.Coordinate, t int) bool {
			return c.Y == t
		}
	}

	filtered := filterObstacles(obstacles, predicate, filterCmp)
	mapped := mapObstacles(filtered, !isX)

	if obstacleHasBiggerIndex {
		for _, item := range mapped {
			if item > searchCmp {
				if isX {
					return utils.Coordinate{X: start.X, Y: item}, true
				} else {
					return utils.Coordinate{X: item, Y: start.Y}, true
				}
			}
		}
	} else {
		for i := len(mapped) - 1; i >= 0; i-- {
			if mapped[i] < searchCmp {
				if isX {
					return utils.Coordinate{X: start.X, Y: mapped[i]}, true
				} else {
					return utils.Coordinate{X: mapped[i], Y: start.Y}, true
				}
			}
		}
	}

	edge := 130
	if isX {
		return utils.Coordinate{X: start.X, Y: edge}, false
	} else {
		return utils.Coordinate{X: edge, Y: start.Y}, false
	}
}

func filterObstacles(obs []utils.Coordinate, predicate func(c utils.Coordinate, cmp int) bool, cmp int) []utils.Coordinate {
	out := make([]utils.Coordinate, 0, len(obs))

	for _, item := range obs {
		if predicate(item, cmp) {
			out = append(out, item)
		}
	}

	return out
}

func mapObstacles(obs []utils.Coordinate, isX bool) []int {
	out := make([]int, 0, len(obs))

	for _, item := range obs {
		if isX {
			out = append(out, item.X)
		} else {
			out = append(out, item.Y)
		}
	}

	return out
}

func (d Day6) Part2(lines []string) (int, error) {
	panic("A")
}
