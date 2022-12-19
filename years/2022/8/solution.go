package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/olofenglund96/aoc-go/helpers"
)

type grid struct {
	points [][]*point
	width  int
	height int
}

type direction struct {
	dx int
	dy int
}

var directions = []direction{
	// Up
	{
		dx: 0,
		dy: -1,
	},
	// Down
	{
		dx: 0,
		dy: 1,
	},
	// Left
	{
		dx: -1,
		dy: 0,
	},
	// Right
	{
		dx: 1,
		dy: 0,
	},
}

type index struct {
	x int
	y int
}

type point struct {
	index   index
	val     int
	visible bool
}

func (i index) Move(d direction) index {
	return index{
		x: i.x + d.dx,
		y: i.y + d.dy,
	}
}

func (i index) Equal(oi index) bool {
	return i.x == oi.x && i.y == oi.y
}

func NewGrid(points [][]int) grid {
	g := grid{
		width:  len(points[0]),
		height: len(points),
	}

	pointGrid := make([][]*point, g.height)
	for i := range pointGrid {
		pointGrid[i] = make([]*point, g.width)
	}

	g.points = pointGrid

	for i, row := range points {
		for j, val := range row {
			g.points[i][j] = &point{
				index: index{
					x: j,
					y: i,
				},
				val: val,
			}
		}
	}

	return g
}

func (g grid) GetPoint(i index) *point {
	return g.points[i.y][i.x]
}

func (g grid) PrintWithPoints(start point, comp point) {
	c := color.New(color.FgCyan, color.Bold)
	r := color.New(color.FgRed, color.Bold)
	gr := color.New(color.FgGreen, color.Bold)
	fmt.Println()
	for _, row := range g.points {
		for _, p := range row {
			if p.index.Equal(start.index) {
				r.Print(p.val)
			} else if p.index.Equal(comp.index) {
				gr.Print(p.val)
			} else if p.visible {
				c.Print(p.val)
			} else {
				fmt.Printf(fmt.Sprint(p.val))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g grid) Print() {
	c := color.New(color.FgCyan, color.Bold)
	fmt.Println()
	for _, row := range g.points {
		for _, p := range row {
			if p.visible {
				c.Print(p.val)
			} else {
				fmt.Printf(fmt.Sprint(p.val))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Print(intGrid [][]int) {
	for _, row := range intGrid {
		fmt.Println(strings.Join(helpers.IntSliceToStrSlice(row), ""))
	}
	fmt.Println()
}

func (g grid) InGrid(i index) bool {
	return i.x >= 0 && i.x < g.width && i.y >= 0 && i.y < g.height
}

func (g grid) VisibleFromDirection(startingPoint point, direction direction) bool {
	currPoint := startingPoint
	for true {
		//helpers.Println("point", currPoint)
		ix := currPoint.index.Move(direction)
		if !g.InGrid(ix) {
			return true
		}

		currPoint = *g.GetPoint(ix)
		//g.PrintWithPoints(startingPoint, currPoint)
		//helpers.WaitForInput()
		if currPoint.val >= startingPoint.val {
			return false
		}
	}

	return true
}

func (g grid) CountVisibleInDirection(startingPoint point, direction direction) int {
	currPoint := startingPoint
	seen := 0
	for true {
		//helpers.Println("point", currPoint)
		ix := currPoint.index.Move(direction)
		if !g.InGrid(ix) {
			return seen
		}

		seen += 1

		currPoint = *g.GetPoint(ix)
		//g.PrintWithPoints(startingPoint, currPoint)
		//helpers.WaitForInput()
		if currPoint.val >= startingPoint.val {
			return seen
		}
	}

	return seen
}

func (g grid) TreeVisible(p point) bool {
	for _, d := range directions {
		//helpers.Println("direction", d)
		if g.VisibleFromDirection(p, d) {
			return true
		}
	}

	return false
}

func (g grid) NumTreesVisibleMul(p point) int {
	totTreesVisible := 1
	for _, d := range directions {
		//helpers.Println("direction", d)
		visibleInDirection := g.CountVisibleInDirection(p, d)
		//helpers.Println("visibleInDirection: ", visibleInDirection)
		totTreesVisible *= visibleInDirection
		//helpers.Println("totTreesVisible: ", totTreesVisible)

	}

	return totTreesVisible
}

func sol1(points [][]int) string {
	treeGrid := NewGrid(points)
	//treeGrid.Print()

	visibleSum := 0
	for _, row := range treeGrid.points {
		for _, p := range row {
			if treeGrid.TreeVisible(*p) {
				p.visible = true
				visibleSum += 1
			}
		}
	}

	//treeGrid.Print()

	return fmt.Sprint(visibleSum)
}

func sol2(points [][]int) string {
	treeGrid := NewGrid(points)
	treeGrid.Print()

	maxVisibleMul := 0
	for _, row := range treeGrid.points {
		for _, p := range row {
			thisMul := treeGrid.NumTreesVisibleMul(*p)
			if thisMul > maxVisibleMul {
				helpers.Println("newMax: ", thisMul, ", point: ", p, ", prevMax: ", maxVisibleMul)
				maxVisibleMul = thisMul
			}

		}
	}

	//treeGrid.Print()

	return fmt.Sprint(maxVisibleMul)
}

func main() {
	rows := helpers.ReadGridFromFile(fmt.Sprintf("years/2022/8/%s.dat", os.Args[2]), "")

	solutionNumer := os.Args[1]
	if solutionNumer == "1" {
		fmt.Print(sol1(rows))
	} else {
		fmt.Print(sol2(rows))
	}
}
