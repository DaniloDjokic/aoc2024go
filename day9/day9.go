package day9

import (
	"aoc2024go/utils"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"strconv"
)

type Day9 struct{}

func (d Day9) Id() int {
	return 9
}

func (d Day9) LoadFile(isTest bool) []string {
	return utils.LoadFile(isTest, d.Id())
}

type DiskElement struct {
	IsEmpty bool
	Id      int
}

type EmptySpace struct {
	Start, Len int
}

func (d Day9) Part1(lines []string) (int, error) {
	disk := make([]DiskElement, 0, 1024)
	emptySpaces := make([]EmptySpace, 0)
	s := stack.New()

	currIdx := 0
	currId := 0

	isFree := false
	line := []rune(lines[0])

	for _, symbol := range line {
		symbolCount, _ := strconv.Atoi(string(symbol))
		end := currIdx + symbolCount

		el := DiskElement{}
		if isFree {
			el.IsEmpty = true
			emptySpaces = append(emptySpaces, EmptySpace{Start: currIdx, Len: end - currIdx})
		} else {
			el.Id = currId
		}

		subSlice := createSubslice(currIdx, end, el, s)
		disk = append(disk, subSlice...)
		currIdx = end

		if !isFree {
			currId++
		}

		isFree = !isFree
	}

	printDisk(disk)

	newLen := len(disk)
	for _, el := range emptySpaces {
		for i := 0; i < el.Len; i++ {
			nextEl := s.Pop().(DiskElement)
			nextIdx := el.Start + i

			disk[nextIdx] = DiskElement{IsEmpty: false, Id: nextEl.Id}
			newLen--
		}
	}

	printDisk(disk)

	total := 0

	for i := 0; i < newLen; i++ {
		total += i * disk[i].Id
	}

	return total, nil
}

func printDisk(disk []DiskElement) {
	for _, el := range disk {
		if el.IsEmpty {
			fmt.Print(string('.'))
		} else {
			fmt.Print(el.Id)
		}
	}

	fmt.Println()
}

func createSubslice(from, to int, val DiskElement, s *stack.Stack) []DiskElement {
	newSlice := make([]DiskElement, to-from)

	for i := 0; i < len(newSlice); i++ {
		newSlice[i] = val
		if !val.IsEmpty {
			s.Push(val)
		}
	}

	return newSlice
}

func (d Day9) Part2(lines []string) (int, error) { panic("AAAA not implemented") }
