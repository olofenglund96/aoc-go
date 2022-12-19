package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/olofenglund96/aoc-go/helpers"
)

func generateDirections(d helpers.Direction, steps int) []helpers.Direction {
	directions := []helpers.Direction{}
	for i := 0; i < steps; i++ {
		directions = append(directions, d)
	}

	return directions
}

func parseInstruction(row string) []helpers.Direction {
	instr := strings.Split(row, " ")
	dir := instr[0]
	steps := helpers.StrToI(instr[1])

	if dir == "R" {
		return generateDirections(helpers.Direction{
			Dx: 1,
			Dy: 0,
		}, steps)
	} else if dir == "U" {
		return generateDirections(helpers.Direction{
			Dx: 0,
			Dy: -1,
		}, steps)
	} else if dir == "L" {
		return generateDirections(helpers.Direction{
			Dx: -1,
			Dy: 0,
		}, steps)
	} else if dir == "D" {
		return generateDirections(helpers.Direction{
			Dx: 0,
			Dy: 1,
		}, steps)
	}

	panic("Could not parse direction")
}

func moveTail(fromHeadPos helpers.Index, toHeadPos helpers.Index, tailPos helpers.Index) helpers.Index {
	if tailPos.Equal(toHeadPos) || tailPos.Neighbour(toHeadPos) || tailPos.DiagNeighbour(toHeadPos) {
		return tailPos
	}

	return fromHeadPos
}

func printGrid(tailIndices map[helpers.Index]bool, headIndex helpers.Index) {
	minX := 1000
	maxX := -1000
	minY := 1000
	maxY := -1000

	tailIndices[headIndex] = false

	for k := range tailIndices {
		if k.X < minX {
			minX = k.X
		}

		if k.X > maxX {
			maxX = k.X
		}

		if k.Y < minY {
			minY = k.Y
		}

		if k.Y > maxY {
			maxY = k.Y
		}
	}

	width := maxX - minX
	height := maxY - minY
	for i := 0; i < height+1; i++ {
		for j := 0; j < width+1; j++ {
			ix := helpers.Index{
				X: j + minX,
				Y: i + minY,
			}

			val, ok := tailIndices[ix]

			if ok && !val {
				fmt.Print("H")
			} else if ok && val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	delete(tailIndices, headIndex)
}

func printGridAll(tailIndices map[helpers.Index]bool, headIndices []helpers.Index) {
	minX := 1000
	maxX := -1000
	minY := 1000
	maxY := -1000

	newCountedIndices := map[helpers.Index]int{}
	for i, v := range headIndices {
		newCountedIndices[v] = i
	}

	helpers.Println("headIndices: ", headIndices)

	for _, k := range headIndices {
		if k.X < minX {
			minX = k.X
		}

		if k.X > maxX {
			maxX = k.X
		}

		if k.Y < minY {
			minY = k.Y
		}

		if k.Y > maxY {
			maxY = k.Y
		}
	}

	width := 40
	height := 40
	fmt.Println()
	for i := 0; i < height+1; i++ {
		for j := 0; j < width+1; j++ {
			ix := helpers.Index{
				X: j - 20,
				Y: i - 20,
			}

			val, ok := newCountedIndices[ix]

			if ok {
				fmt.Print(fmt.Sprint(val))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func sol1(rows []string) string {
	tailIndices := map[helpers.Index]bool{}
	headIndex := helpers.Index{
		X: 0,
		Y: 0,
	}
	tailIndex := headIndex
	tailIndices[tailIndex] = true

	for _, instr := range rows {
		directions := parseInstruction(instr)
		//helpers.Println("instruction: ", instr, ", directions: ", directions)
		for _, d := range directions {
			newHeadIndex := headIndex.Move(d)
			newTailIndex := moveTail(headIndex, newHeadIndex, tailIndex)
			if _, ok := tailIndices[newTailIndex]; !ok {
				helpers.Println("Moved from: ", tailIndex, ", to: ", newTailIndex)
			} else {
				helpers.Println("Stayed at: ", tailIndex)
			}
			tailIndices[newTailIndex] = true

			headIndex = newHeadIndex
			tailIndex = newTailIndex
			//printGrid(tailIndices, headIndex)
			//helpers.WaitForInput()
		}
	}

	//printGrid(tailIndices)

	//helpers.Println("tailIndices: ", tailIndices)

	return fmt.Sprint(len(tailIndices))
}

func moveTailv2(fromHeadPos helpers.Index, toHeadPos helpers.Index, tailPos helpers.Index) helpers.Index {
	if tailPos.Equal(toHeadPos) || tailPos.Neighbour(toHeadPos) || tailPos.DiagNeighbour(toHeadPos) {
		return tailPos
	}

	if tailPos.EqualX(toHeadPos) {
		return helpers.Index{
			X: tailPos.X,
			Y: toHeadPos.Y - (toHeadPos.Y-tailPos.Y)/int(math.Abs(float64((toHeadPos.Y-tailPos.Y)))),
		}
	}

	if tailPos.EqualY(toHeadPos) {
		return helpers.Index{
			Y: tailPos.Y,
			X: toHeadPos.X - (toHeadPos.X-tailPos.X)/int(math.Abs(float64((toHeadPos.X-tailPos.X)))),
		}
	}

	if tailPos.DiagNeighbour(fromHeadPos) {
		direction := helpers.DirectionFromIndices(fromHeadPos, toHeadPos)
		return toHeadPos.Move(helpers.Direction{
			Dx: -direction.Dx,
			Dy: -direction.Dy,
		})
	}

	direction := helpers.DirectionFromIndices(tailPos, fromHeadPos)

	return toHeadPos.Move(helpers.Direction{
		Dx: -direction.Dx,
		Dy: -direction.Dy,
	})
}

func sol2(rows []string) string {
	tailIndices := map[helpers.Index]bool{}
	headIndices := []helpers.Index{}
	for i := 0; i < 10; i++ {
		headIndices = append(headIndices, helpers.Index{
			X: 0,
			Y: 0,
		})
	}

	tailIndex := headIndices[len(headIndices)-1]
	tailIndices[tailIndex] = true

	for _, instr := range rows {
		directions := parseInstruction(instr)
		//helpers.Println("instruction: ", instr, ", directions: ", directions)
		for _, d := range directions {
			newHeadIndices := make([]helpers.Index, len(headIndices))
			headIndex := headIndices[0]
			newHeadIndices[0] = headIndex.Move(d)
			//helpers.Println("Moved ", 0, " from: ", headIndex, ", to: ", newHeadIndices[0])

			for i, ix := range headIndices[1:] {
				newPos := moveTailv2(headIndices[i], newHeadIndices[i], ix)
				helpers.Println("Moved ", i+1, " from: ", ix, ", to: ", newPos)
				newHeadIndices[i+1] = newPos
			}

			newTailIndex := newHeadIndices[len(newHeadIndices)-1]

			if _, ok := tailIndices[newTailIndex]; !ok {
				helpers.Println("Moved from: ", tailIndex, ", to: ", newTailIndex)
			} else {
				helpers.Println("Stayed at: ", tailIndex)
			}
			tailIndices[newTailIndex] = true
			headIndices = newHeadIndices

			//printGridAll(tailIndices, headIndices)
			//helpers.WaitForInput()
		}
	}

	printGrid(tailIndices, helpers.Index{X: 0, Y: 0})

	helpers.Println("tailIndices: ", tailIndices)

	return fmt.Sprint(len(tailIndices) + 1)
}

func main() {
	rows := helpers.ReadFileLines(fmt.Sprintf("years/2022/9/%s.dat", os.Args[2]))

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
