package helpers

import (
	"os"
	"strings"
)

func ReadFileLines(filePath string) []string {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	dat = dat[:len(dat)-1]

	return strings.Split(string(dat), "\n")
}

func ReadGridFromFile(filePath string, sep string) [][]int {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	dat = dat[:len(dat)-1]

	var grid [][]int

	for _, row := range strings.Split(string(dat), "\n") {
		chars := strings.Split(row, sep)
		grid = append(grid, StrSliceToIntSlice(chars))
	}

	return grid
}

func ReadGridFromFileWithFunc[K comparable](filePath string, parseFunc func(index Index, c rune) Cell[K]) [][]*Cell[K] {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	dat = dat[:len(dat)-1]

	var grid [][]*Cell[K]

	for i, row := range strings.Split(string(dat), "\n") {
		var items []*Cell[K]
		for j, c := range row {
			cell := parseFunc(Index{X: j, Y: i}, c)
			items = append(items, &cell)
		}

		grid = append(grid, items)
	}

	return grid
}
