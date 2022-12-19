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

func ReadGridFromFileWithFunc(filePath string, parseFunc func(c rune) interface{}) [][]interface{} {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	dat = dat[:len(dat)-1]

	var grid [][]interface{}

	for _, row := range strings.Split(string(dat), "\n") {
		var ints []interface{}
		for _, c := range row {
			ints = append(ints, parseFunc(c))
		}

		grid = append(grid, ints)
	}

	return grid
}
