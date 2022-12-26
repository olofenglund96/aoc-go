package main

import (
	"fmt"
	"os"

	"github.com/olofenglund96/aoc-go/helpers"
)

type Stack[C any] struct {
	items []C
}

func (s *Stack[C]) getItemIndexWithLowestPriority(prioFun func(c C) int) (C, int) {
	lowestItem := s.items[0]
	lowestIx := 0
	for ix, item := range s.items[1:] {
		if prioFun(lowestItem) > prioFun(item) {
			lowestItem = item
			lowestIx = ix
		}
	}

	return lowestItem, lowestIx
}

func (s *Stack[C]) Pop(prioFun func(c C) int) C {
	i, ix := s.getItemIndexWithLowestPriority(prioFun)
	items := s.items
	iB := items[:ix]
	iA := items[ix+1:]
	s.items = append(iB, iA...)
	return i
}

func (s *Stack[C]) Push(items ...C) {
	s.items = append(s.items, items...)
}

func (s *Stack[C]) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack[C]) String() string {
	prtStr := ""
	for _, i := range s.items {
		prtStr += fmt.Sprintf("%v, ", i)
	}

	return prtStr
}

func getNeighbours(grid helpers.Grid[graphPoint], cell *helpers.Cell[graphPoint]) []*helpers.Cell[graphPoint] {
	neighbourIxs := []helpers.Index{}
	for _, i := range cell.Index.GetManhattanNeighbours() {
		if grid.InGrid(i) {
			neighbourIxs = append(neighbourIxs, i)
		}
	}

	accessibleNeighbours := []*helpers.Cell[graphPoint]{}
	for _, ix := range neighbourIxs {
		c := grid.GetCell(ix)
		if !c.Marked && grid.IsConnected(*cell, *c) {
			c.Marked = true
			c.Context.distanceToStart = cell.Context.distanceToStart + 1
			accessibleNeighbours = append(accessibleNeighbours, c)
		}
	}

	return accessibleNeighbours
}

func prioFun(cell *helpers.Cell[graphPoint]) int {
	return cell.Context.distanceToStart
}

func bfs(grid helpers.Grid[graphPoint], start *helpers.Cell[graphPoint], end *helpers.Cell[graphPoint]) int {
	start.Marked = true
	pointStack := Stack[*helpers.Cell[graphPoint]]{
		items: []*helpers.Cell[graphPoint]{start},
	}

	for !pointStack.Empty() {
		//helpers.Println("PointStack bf: ", pointStack.String())
		p := pointStack.Pop(prioFun)
		//helpers.Println("PointStack af: ", pointStack.String())
		//helpers.Println("In point: ", p)
		if p.Repr == "E" {
			helpers.Println(p.Repr)
			fmt.Println("DONE!")
			//grid.PrintBeautifully(p)
			return p.Context.distanceToStart
		}
		neighbours := getNeighbours(grid, p)
		pointStack.Push(neighbours...)
		//helpers.Println("Neighbours: ", neighbours)
		//helpers.Println("PointStack: ", pointStack.String())
		//grid.PrintBeautifully(p)
		//helpers.WaitForInput()
	}

	//grid.PrintBeautifully()

	return 100000
}

func printContext(c helpers.Cell[graphPoint]) string {
	return c.Context.String() + "\t"
}

// Example solution: 31
func sol1(grid helpers.Grid[graphPoint]) string {
	//grid.Print()
	//grid.PrintContext(printContext)

	var start *helpers.Cell[graphPoint]
	var end *helpers.Cell[graphPoint]
	for _, row := range grid.Cells {
		for _, c := range row {
			if c.Repr == "S" {
				start = c
			}
			if c.Repr == "E" {
				end = c
			}
		}
	}

	return fmt.Sprint(bfs(grid, start, end))
}

func resetContextFun(context graphPoint) graphPoint {
	return graphPoint{
		value:           context.value,
		distanceToStart: 0,
	}
}

func sol2(grid helpers.Grid[graphPoint]) string {
	var starts []*helpers.Cell[graphPoint]
	var end *helpers.Cell[graphPoint]
	for _, row := range grid.Cells {
		for _, c := range row {
			if c.Context.value == 64 {
				starts = append(starts, c)
			}
			if c.Repr == "E" {
				end = c
			}
		}
	}

	minSteps := 10000
	for _, s := range starts {
		steps := bfs(grid, s, end)
		if steps < minSteps {
			minSteps = steps
		}

		grid.Reset(resetContextFun)
	}

	return fmt.Sprint(minSteps)
}

type graphPoint struct {
	distanceToStart int
	value           int
}

func (g graphPoint) String() string {
	return fmt.Sprint(g.value)
}

func (g graphPoint) Priority() int {
	return g.distanceToStart
}

func accessible(c1 helpers.Cell[graphPoint], c2 helpers.Cell[graphPoint]) bool {
	return c1.Context.value >= c2.Context.value-1
}

func readFunc[K comparable](index helpers.Index, c rune) helpers.Cell[graphPoint] {
	repr := string(c)
	if c == 'S' {
		c = 'a'
	} else if c == 'E' {
		c = 'z'
	}

	intVal := int(c)
	p := helpers.Cell[graphPoint]{
		Index: index,
		Repr:  repr,
		Context: graphPoint{
			value:           intVal - 65 + (97 - 65),
			distanceToStart: 0,
		},
		Marked: false,
	}

	return p
}

func main() {
	iGrid := helpers.ReadGridFromFileWithFunc(fmt.Sprintf("years/2022/12/%s.dat", os.Args[2]), readFunc[graphPoint])
	grid := helpers.NewGridWithConnectFun(iGrid, accessible)
	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(grid))
	} else {
		fmt.Print(sol2(grid))
	}
}
