package main

import (
	"fmt"
	"os"

	"github.com/olofenglund96/aoc-go/helpers"
)

type gridVal struct {
	char  string
	value int
}

func (g gridVal) String() string {
	return fmt.Sprint(g.char)
}

func (g gridVal) accessible(gTo gridVal) bool {
	return g.value >= gTo.value-1 && g.value <= gTo.value+1
}

func Pop(s []*helpers.Point) ([]*helpers.Point, *helpers.Point) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func getNeighbours(grid helpers.Grid, point *helpers.Point) []*helpers.Point {
	neighbourIxs := point.Index.GetManhattanNeighbours()
	accessibleNeighbours := []*helpers.Point{}
	for _, ix := range neighbourIxs {
		p := grid.GetPoint(ix)
		if !p.Marked && point.Val.(gridVal).accessible(p.Val.(gridVal)) {
			p.Marked = true
			accessibleNeighbours = append(accessibleNeighbours, p)
		}
	}

	return accessibleNeighbours
}

func bfs(grid helpers.Grid, start *helpers.Point, end *helpers.Point) {
	interestingPoints := []*helpers.Point{start}

	for len(interestingPoints) > 0 {
		interestingPoints, p := Pop(interestingPoints)
		neighbours := getNeighbours(grid, p)
		helpers.Println(interestingPoints, neighbours)
	}
}

func sol1(grid helpers.Grid) string {
	grid.Print()
	bfs(grid, grid.Points[0][0], grid.Points[1][1])
	return "Solution1"
}

func sol2(rows helpers.Grid) string {

	return "Solution2"
}

func readFunc(c rune) interface{} {
	intVal := int(c)
	if intVal >= 97 {
		return gridVal{
			char:  string(c),
			value: intVal - 97,
		}
	}

	return gridVal{
		char:  string(c),
		value: intVal - 65 + (97 - 65),
	}
}

func main() {
	iGrid := helpers.ReadGridFromFileWithFunc(fmt.Sprintf("years/2022/12/%s.dat", os.Args[2]), readFunc)
	grid := helpers.NewGrid(iGrid)
	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(grid))
	} else {
		fmt.Print(sol2(grid))
	}
}
