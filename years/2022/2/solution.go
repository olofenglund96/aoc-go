package main

import (
	"fmt"
	"os"

	"strings"
)

var winMap = map[string]string{
	"X": "C",
	"Y": "A",
	"Z": "B",
}

var winMapRev = map[string]string{
	"C": "X",
	"A": "Y",
	"B": "Z",
}

var drawMap = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var drawMapRev = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var looseMapRev = map[string]string{
	"B": "X",
	"C": "Y",
	"A": "Z",
}

var scoreMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

const (
	winScore  = 6
	drawScore = 3
)

func sol1(input string) string {
	rows := strings.Split(input, "\n")
	score := 0
	for _, row := range rows {
		plays := strings.Split(row, " ")
		theirPlay := plays[0]

		myPlay := plays[1]
		score += scoreMap[myPlay]

		if val, ok := winMap[myPlay]; ok && val == theirPlay {
			score += winScore
		} else if val, ok := drawMap[myPlay]; ok && val == theirPlay {
			score += drawScore
		}

		//fmt.Printf("Row %d: Score=%d\n", i, score)
	}

	return fmt.Sprint(score)
}

func sol2(input string) string {
	rows := strings.Split(input, "\n")
	score := 0
	for _, row := range rows {
		plays := strings.Split(row, " ")
		theirPlay := plays[0]
		myPlay := plays[1]
		myActualPlay := looseMapRev[theirPlay]

		if myPlay == "Z" {
			myActualPlay = winMapRev[theirPlay]
			score += winScore
		} else if myPlay == "Y" {
			myActualPlay = drawMapRev[theirPlay]
			score += drawScore
		}

		score += scoreMap[myActualPlay]

		//fmt.Printf("Row %d: Score=%d\n", i, score)
	}

	return fmt.Sprint(score)
}

func main() {
	dat, err := os.ReadFile(fmt.Sprintf("years/2022/2/%s.dat", os.Args[2]))
	dat = dat[:len(dat)-1]

	if err != nil {
		panic(err)
	}
	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(string(dat)))
	} else {
		fmt.Print(sol2(string(dat)))
	}
}
