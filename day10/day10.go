package day10

import (
	"aoc2024go/utils"
	"strconv"
	"sync"
)

type Day10 struct{}

func (d Day10) Id() int {
	return 10
}

func (d Day10) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

func (d Day10) Part1(lines []string) (int, error) {
	matrix := utils.ConvertToCharArray(lines)
	total := 0
	trails := make(map[utils.Coordinate]map[utils.Coordinate]bool)
	lock := sync.Mutex{}
	dimX := len(matrix)
	dimY := len(matrix[0])

	wg := sync.WaitGroup{}

	for i := range matrix {
		for j, item := range matrix[i] {
			cellNumber, _ := strconv.Atoi(string(item))

			if cellNumber == 0 {
				peaks := make(map[utils.Coordinate]bool)
				start := utils.Coordinate{X: i, Y: j}
				wg.Add(1)
				go func() {
					defer wg.Done()
					findPeaks(start, peaks, &lock, dimX, dimY, matrix)
				}()
				trails[start] = peaks
			}
		}
	}

	wg.Wait()

	for _, v := range trails {
		peakCount := len(v)
		total += peakCount
	}

	return total, nil
}

func findPeaks(curr utils.Coordinate, peaks map[utils.Coordinate]bool, lock *sync.Mutex, dimX, dimY int, matrix [][]rune) {
	currNum, _ := strconv.Atoi(string(matrix[curr.X][curr.Y]))

	if currNum == 9 {
		lock.Lock()
		if _, ok := peaks[curr]; !ok {
			peaks[curr] = true
		}
		lock.Unlock()
	}

	left := utils.Coordinate{X: curr.X, Y: curr.Y - 1}
	tryRecurse(left, currNum, peaks, lock, dimX, dimY, matrix)
	right := utils.Coordinate{X: curr.X, Y: curr.Y + 1}
	tryRecurse(right, currNum, peaks, lock, dimX, dimY, matrix)
	top := utils.Coordinate{X: curr.X - 1, Y: curr.Y}
	tryRecurse(top, currNum, peaks, lock, dimX, dimY, matrix)
	bot := utils.Coordinate{X: curr.X + 1, Y: curr.Y}
	tryRecurse(bot, currNum, peaks, lock, dimX, dimY, matrix)
}

func tryRecurse(curr utils.Coordinate, currNum int, total map[utils.Coordinate]bool, lock *sync.Mutex, dimX, dimY int, matrix [][]rune) {
	if !utils.IsOffMap(curr, dimX, dimY) {
		leftNum, _ := strconv.Atoi(string(matrix[curr.X][curr.Y]))
		if leftNum-currNum == 1 {
			findPeaks(curr, total, lock, dimX, dimY, matrix)
		}
	}
}

func (d Day10) Part2(lines []string) (int, error) {
	panic("AA")
}
