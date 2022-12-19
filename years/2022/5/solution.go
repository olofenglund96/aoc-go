package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olofenglund96/aoc-go/helpers"
)

func Pop(s []rune) ([]rune, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

const colWidth = 4

type moveInstruction struct {
	numCrates int
	fromCol   int
	toCol     int
}

func parseInput(rows []string) ([][]rune, []moveInstruction) {
	var crateConfig []string
	var moveProcedure []string
	doneCrateConfig := false
	for _, r := range rows {
		if r == "" {
			doneCrateConfig = true
			continue
		}
		if !doneCrateConfig {
			crateConfig = append(crateConfig, r)
		} else {
			moveProcedure = append(moveProcedure, r)
		}
	}

	stacks := make([][]rune, len(crateConfig[0])/colWidth+1)
	crateConfig = crateConfig[:len(crateConfig)-1]
	for _, r := range crateConfig {
		for cIx, c := range r {
			if c != '[' && c != ']' && c != ' ' {
				col := cIx / colWidth

				stacks[col] = append([]rune{c}, stacks[col]...)
			}
		}
	}

	var moveInstructions []moveInstruction
	for _, mp := range moveProcedure {
		words := strings.Split(mp, " ")
		numPop := helpers.StrToI(words[1])
		fromCol := helpers.StrToI(words[3])
		toCol := helpers.StrToI(words[5])
		moveInstructions = append(moveInstructions, moveInstruction{
			numCrates: numPop,
			fromCol:   fromCol - 1,
			toCol:     toCol - 1,
		})
	}

	return stacks, moveInstructions
}

func moveCrates(stack [][]rune, instr moveInstruction) [][]rune {
	fmt.Println(instr)
	for i := 0; i < instr.numCrates; i += 1 {
		println(fmt.Sprintf("stacks: %v", stack))
		colStack, crate := Pop(stack[instr.fromCol])
		stack[instr.fromCol] = colStack

		stack[instr.toCol] = append(stack[instr.toCol], crate)
	}

	return stack
}

func sol1(rows []string) string {
	stacks, moveProc := parseInput(rows)
	println(fmt.Sprintf("stacks: %+v", stacks))
	fmt.Println(moveProc)

	for _, mp := range moveProc {
		stacks = moveCrates(stacks, mp)
	}

	resString := ""
	for _, s := range stacks {
		resString = resString + string(s[len(s)-1])
	}

	return resString
}

func moveCratesKeepOrder(stack [][]rune, instr moveInstruction) [][]rune {
	fmt.Println(instr)
	var cratesMoving []rune
	for i := 0; i < instr.numCrates; i += 1 {
		println(fmt.Sprintf("stacks: %v", stack))
		colStack, crate := Pop(stack[instr.fromCol])
		stack[instr.fromCol] = colStack

		cratesMoving = append([]rune{crate}, cratesMoving...)
	}

	for _, crate := range cratesMoving {
		stack[instr.toCol] = append(stack[instr.toCol], crate)
	}

	return stack
}

func sol2(rows []string) string {
	stacks, moveProc := parseInput(rows)
	helpers.Println("stacks: ", stacks, ", moveProc: ", moveProc)

	for _, mp := range moveProc {
		stacks = moveCratesKeepOrder(stacks, mp)
	}

	resString := ""
	for _, s := range stacks {
		resString = resString + string(s[len(s)-1])
	}

	return resString
}

func main() {
	rows := helpers.ReadFileLines(fmt.Sprintf("years/2022/5/%s.dat", os.Args[2]))

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
