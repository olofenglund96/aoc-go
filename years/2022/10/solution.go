package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/olofenglund96/aoc-go/helpers"
)

func parseAddxInstr(inp string) int {
	inpS := strings.Split(inp, " ")
	return helpers.StrToI(inpS[1])
}

func interestingCycle(cycle int) bool {
	return (cycle+20)%40 == 0
}

func sol1(rows []string) string {
	//rows = []string{"noop", "addx 3", "addx -5"}
	reg := 1
	lineIx := 0
	cycle := 1
	regSum := 0
	numCyclesToSkip := 0
	for lineIx < len(rows) {
		helpers.Println("[Cycle start ", cycle, "] instr: ", rows[lineIx], ", regSum: ", regSum, ", reg: ", reg)
		//helpers.WaitForInput()
		if interestingCycle(cycle) {
			helpers.Println("[Interesting cycle ", cycle, "] regSum: ", regSum, ", reg: ", reg)
			regSum += reg * cycle
		}

		cycle++

		if numCyclesToSkip > 0 {
			numCyclesToSkip--

			if numCyclesToSkip == 0 {
				reg += parseAddxInstr(rows[lineIx])
				lineIx++
			}
			continue
		}

		currRow := rows[lineIx]

		if strings.Contains(currRow, "addx") {
			numCyclesToSkip = 1
			continue
		}

		lineIx++
	}

	//helpers.Println("[Cycle end] regSum: ", regSum, ", reg: ", reg)

	return fmt.Sprint(regSum)
}

func IMin(x int, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func IMax(x int, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func sol2(rows []string) string {

	intGrid := make([][]int, 6)
	for i := range intGrid {
		intGrid[i] = make([]int, 40)
	}

	grid := helpers.NewGrid(intGrid)
	grid.Print()
	reg := 1
	lineIx := 0
	cycle := 1
	regSum := 0
	numCyclesToSkip := 0
	for lineIx < len(rows) {
		helpers.Println("[Cycle start ", cycle, "] instr: ", rows[lineIx], ", reg: ", reg, " cycleMod: ", (cycle-1)%40)

		for i := reg - 1; i <= reg+1; i++ {
			if i == (cycle-1)%40 {
				iRow := int(math.Floor(float64(cycle / 40)))
				iCol := (cycle - 1) % 40

				//helpers.Println("Matching sprite-cycle [", cycle, "] row: ", iRow, ", col: ", iCol)

				grid.Points[iRow][iCol] = &helpers.Point{
					Val:    1,
					Marked: true,
				}

				grid.Print()
				//helpers.WaitForInput()
			}
		}

		cycle++

		if numCyclesToSkip > 0 {
			numCyclesToSkip--

			if numCyclesToSkip == 0 {
				reg += parseAddxInstr(rows[lineIx])
				lineIx++
			}
			continue
		}

		currRow := rows[lineIx]

		if strings.Contains(currRow, "addx") {
			numCyclesToSkip = 1
			continue
		}

		lineIx++
	}

	grid.PrintMarked()

	//helpers.Println("[Cycle end] regSum: ", regSum, ", reg: ", reg)

	return fmt.Sprint(regSum)
}

func main() {
	rows := helpers.ReadFileLines(fmt.Sprintf("years/2022/10/%s.dat", os.Args[2]))

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
