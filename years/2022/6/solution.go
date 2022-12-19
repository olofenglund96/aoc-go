package main

import (
	"fmt"
	"os"

	"github.com/olofenglund96/aoc-go/helpers"
)

func padString(str string, padStr string, padLen int) string {
	newStr := str
	for i := 0; i < padLen; i++ {
		newStr = padStr + newStr + padStr
	}

	return newStr
}

func AddOrIncr(countMap map[rune]int, newRune rune) {
	_, exists := countMap[newRune]
	if exists {
		countMap[newRune] += 1
	} else {
		countMap[newRune] = 1
	}
}

func RemoveOrDecr(countMap map[rune]int, newRune rune) {
	count := countMap[newRune]
	if count == 1 {
		delete(countMap, newRune)
	} else {
		countMap[newRune] -= 1
	}
}

func UniqueRunes(countMap map[rune]int) bool {
	for _, c := range countMap {
		if c != 1 {
			return false
		}
	}

	return true
}

func GetKeyString(countMap map[rune]int) string {
	keyString := ""
	for k, _ := range countMap {
		keyString += string(k)
	}

	return keyString
}

func PrintMap(countMap map[rune]int) {
	for k, v := range countMap {
		fmt.Printf("%s -> %d\n", string(k), v)
	}
}

func sol1(rows []string) string {
	subroutine := rows[0]
	helpers.Println("example input: ", subroutine)
	//subrPadded := padString(subroutine, "_", 4)

	i := 4
	runeCounts := map[rune]int{}
	for _, c := range subroutine[:i] {
		AddOrIncr(runeCounts, c)
	}

	for i < len(subroutine) {
		fmt.Printf("Iteration %d, subroutine %s\n", i, string(subroutine[i]))
		PrintMap(runeCounts)
		if UniqueRunes(runeCounts) {
			helpers.Println(GetKeyString(runeCounts))
			return fmt.Sprint(i)
		}
		RemoveOrDecr(runeCounts, rune(subroutine[i-4]))
		AddOrIncr(runeCounts, rune(subroutine[i]))
		i++
	}

	return "No solution foudnd"
}

func sol2(rows []string) string {
	subroutine := rows[0]
	helpers.Println("example input: ", subroutine)
	//subrPadded := padString(subroutine, "_", 4)

	i := 14
	runeCounts := map[rune]int{}
	for _, c := range subroutine[:i] {
		AddOrIncr(runeCounts, c)
	}

	for i < len(subroutine) {
		fmt.Printf("Iteration %d, subroutine %s\n", i, string(subroutine[i]))
		PrintMap(runeCounts)
		if UniqueRunes(runeCounts) {
			helpers.Println(GetKeyString(runeCounts))
			return fmt.Sprint(i)
		}
		RemoveOrDecr(runeCounts, rune(subroutine[i-14]))
		AddOrIncr(runeCounts, rune(subroutine[i]))
		i++
	}

	return "Solution2"
}

func main() {
	rows := helpers.ReadFileLines(fmt.Sprintf("years/2022/6/%s.dat", os.Args[2]))

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
