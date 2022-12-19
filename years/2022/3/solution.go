package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

const asciiOffsetLower = int('a') - 1
const asciiOffsetUpper = int('A') - 27

func findCommonLetter(word1 string, word2 string) (rune, error) {
	for _, l := range word1 {
		if strings.ContainsRune(word2, l) {
			return l, nil
		}
	}

	return 'x', fmt.Errorf("could not find rune in other string")
}

func sol1(input string) string {
	rows := strings.Split(input, "\n")

	score := 0
	for i, row := range rows {
		rowLen := len(row)
		p1 := row[:rowLen/2]
		p2 := row[rowLen/2:]

		letter, err := findCommonLetter(p1, p2)
		if err != nil {
			panic(fmt.Sprintf("Did not find common letter in %s and %s", p1, p2))
		}

		if unicode.IsUpper(letter) {
			fmt.Printf("Upper: %c, val: %d\n", letter, int(letter)-asciiOffsetUpper)
			score += int(letter) - asciiOffsetUpper
		} else {
			fmt.Printf("Lower: %c, val: %d\n", letter, int(letter)-asciiOffsetLower)
			score += int(letter) - asciiOffsetLower
		}

		fmt.Println(fmt.Sprintf("Iteration %d, score: %d", i, score))
	}

	return fmt.Sprint(score)
}

func findCommonLetterThree(word1 string, word2 string, word3 string) (rune, error) {
	for _, l := range word1 {
		if strings.ContainsRune(word2, l) && strings.ContainsRune(word3, l) {
			return l, nil
		}
	}

	return 'x', fmt.Errorf("could not find common rune")
}

func sol2(input string) string {
	rows := strings.Split(input, "\n")

	score := 0
	i := 0
	for i < len(rows) {
		r1 := rows[i]
		r2 := rows[i+1]
		r3 := rows[i+2]

		letter, err := findCommonLetterThree(r1, r2, r3)
		if err != nil {
			panic(err)
		}

		if unicode.IsUpper(letter) {
			fmt.Printf("Upper: %c, val: %d\n", letter, int(letter)-asciiOffsetUpper)
			score += int(letter) - asciiOffsetUpper
		} else {
			fmt.Printf("Lower: %c, val: %d\n", letter, int(letter)-asciiOffsetLower)
			score += int(letter) - asciiOffsetLower
		}

		fmt.Println(fmt.Sprintf("Iteration %d, score: %d", i, score))
		i += 3
	}

	return fmt.Sprint(score)
}

func main() {
	dat, err := os.ReadFile(fmt.Sprintf("years/2022/3/%s.dat", os.Args[2]))
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
