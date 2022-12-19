package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func strSliceToIntSlice(strSlice []string) []int {
	var iSlice []int

	for _, s := range strSlice {
		iSlice = append(iSlice, strToI(s))
	}

	return iSlice
}

func sol1(input []string) string {
	count := 0
	for _, row := range input {
		pairs := strings.Split(row, ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		c1 := (strToI(p1[0]) >= strToI(p2[0]) && strToI(p1[1]) <= strToI(p2[1]))
		c2 := (strToI(p2[0]) >= strToI(p1[0]) && strToI(p2[1]) <= strToI(p1[1]))
		contained := c1 || c2
		fmt.Printf("p1: %v, p2: %v, contained: %v, c1: %v, c2: %v\n", p1, p2, contained, c1, c2)

		if contained {
			count += 1
		}
	}

	return fmt.Sprint(count)
}

func sol2(input []string) string {
	count := 0
	for _, row := range input {
		pairs := strings.Split(row, ",")
		p1 := strSliceToIntSlice(strings.Split(pairs[0], "-"))
		p2 := strSliceToIntSlice(strings.Split(pairs[1], "-"))

		overlap := (p1[0] >= p2[0] && p1[0] <= p2[1]) || (p1[1] >= p2[0] && p1[1] <= p2[1]) || (p2[0] > p1[0] && p2[0] < p1[1])

		fmt.Printf("p1: %v, p2: %v, overlap: %v\n", p1, p2, overlap)

		if overlap {
			count += 1
		}
	}

	return fmt.Sprint(count)
}

func main() {
	dat, err := os.ReadFile(fmt.Sprintf("years/2022/4/%s.dat", os.Args[2]))
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)-1]

	rows := strings.Split(string(dat), "\n")

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
