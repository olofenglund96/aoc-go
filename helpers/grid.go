package helpers

import (
	"fmt"

	"github.com/fatih/color"
)

// GRID
type Grid struct {
	Points [][]*Point
	Width  int
	Height int
}

func NewGrid(points [][]interface{}) Grid {
	g := Grid{
		Width:  len(points[0]),
		Height: len(points),
	}

	pointGrid := make([][]*Point, g.Height)
	for i := range pointGrid {
		pointGrid[i] = make([]*Point, g.Width)
	}

	g.Points = pointGrid

	for i, row := range points {
		for j, val := range row {
			g.Points[i][j] = &Point{
				Index: Index{
					X: j,
					Y: i,
				},
				Val: val,
			}
		}
	}

	return g
}

func (g Grid) GetPoint(i Index) *Point {
	return g.Points[i.Y][i.X]
}

func (g Grid) InGrid(i Index) bool {
	return i.X >= 0 && i.X < g.Width && i.Y >= 0 && i.Y < g.Height
}

func (g Grid) PrintWithPoints(start Point, comp Point) {
	c := color.New(color.FgCyan, color.Bold)
	r := color.New(color.FgRed, color.Bold)
	gr := color.New(color.FgGreen, color.Bold)
	fmt.Println()
	for _, row := range g.Points {
		for _, p := range row {
			if p.Index.Equal(start.Index) {
				r.Print(p.Val)
			} else if p.Index.Equal(comp.Index) {
				gr.Print(p.Val)
			} else if p.Marked {
				c.Print(p.Val)
			} else {
				fmt.Printf(fmt.Sprint(p.Val))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) Print() {
	c := color.New(color.FgCyan, color.Bold)
	fmt.Println()
	for _, row := range g.Points {
		for _, p := range row {
			if p.Marked {
				c.Print(p.Val)
			} else {
				fmt.Printf(fmt.Sprint(p.Val))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) PrintMarked() {
	fmt.Println()
	for _, row := range g.Points {
		for _, p := range row {
			if p.Marked {
				fmt.Printf(fmt.Sprint("#"))
			} else {
				fmt.Printf(fmt.Sprint(" "))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Direction struct {
	Dx int
	Dy int
}

// INDEX
type Index struct {
	X int
	Y int
}

func (i Index) Move(d Direction) Index {
	return Index{
		X: i.X + d.Dx,
		Y: i.Y + d.Dy,
	}
}

func DirectionFromIndices(iFrom Index, iTo Index) Direction {
	return Direction{
		Dx: iTo.X - iFrom.X,
		Dy: iTo.Y - iFrom.Y,
	}
}

func (i Index) EqualX(oi Index) bool {
	return i.X == oi.X
}

func (i Index) EqualY(oi Index) bool {
	return i.Y == oi.Y
}

func (i Index) Equal(oi Index) bool {
	return i.EqualX(oi) && i.EqualY(oi)
}

func (i Index) NeighbourX(oi Index) bool {
	return (i.X-1 == oi.X && i.EqualY(oi)) || (i.X+1 == oi.X && i.EqualY(oi))
}

func (i Index) NeighbourY(oi Index) bool {
	return (i.Y-1 == oi.Y && i.EqualX(oi)) || (i.Y+1 == oi.Y && i.EqualX(oi))
}

func (i Index) Neighbour(oi Index) bool {
	return i.NeighbourX(oi) || i.NeighbourY(oi)
}

func (i Index) NeighbourTL(oi Index) bool {
	return (i.X-1 == oi.X && i.Y-1 == oi.Y)
}

func (i Index) NeighbourTR(oi Index) bool {
	return (i.X+1 == oi.X && i.Y-1 == oi.Y)
}

func (i Index) NeighbourBL(oi Index) bool {
	return (i.X-1 == oi.X && i.Y+1 == oi.Y)
}

func (i Index) NeighbourBR(oi Index) bool {
	return (i.X+1 == oi.X && i.Y+1 == oi.Y)
}

func (i Index) DiagNeighbour(oi Index) bool {
	return i.NeighbourTL(oi) || i.NeighbourTR(oi) || i.NeighbourBL(oi) || i.NeighbourBR(oi)
}

func (i Index) GetManhattanNeighbours() []Index {
	return []Index{
		{
			X: i.X - 1,
			Y: i.Y,
		},
		{
			X: i.X + 1,
			Y: i.Y,
		},
		{
			X: i.X,
			Y: i.Y - 1,
		},
		{
			X: i.X,
			Y: i.Y - 1,
		},
	}
}

type Point struct {
	Index  Index
	Val    interface{}
	Marked bool
}
