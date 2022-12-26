package helpers

import (
	"fmt"

	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

// GRID
type Grid[K comparable] struct {
	Cells      [][]*Cell[K]
	Width      int
	Height     int
	connectFun func(c1 Cell[K], c2 Cell[K]) bool
}

func (g Grid[K]) IsConnected(c1 Cell[K], c2 Cell[K]) bool {
	return g.connectFun(c1, c2)
}

type Cell[K comparable] struct {
	Index   Index
	Context K
	Repr    string
	Marked  bool
}

func (c Cell[K]) String() string {
	return fmt.Sprintf("[%d,%d]->%s", c.Index.X, c.Index.Y, c.Repr)
}

func (c Cell[K]) Equal(cell Cell[K]) bool {
	return c.Index.Equal(cell.Index)
}

func (c Cell[K]) Accessible(cell Cell[K]) bool {
	return c.Index.Equal(cell.Index)
}

func NewGrid[K comparable](cells [][]*Cell[K]) Grid[K] {
	g := Grid[K]{
		Width:  len(cells[0]),
		Height: len(cells),
	}

	cellGrid := make([][]*Cell[K], g.Height)
	for i := range cellGrid {
		cellGrid[i] = make([]*Cell[K], g.Width)
	}

	g.Cells = cells

	return g
}

func NewGridWithConnectFun[K comparable](cells [][]*Cell[K], connectFun func(c1 Cell[K], c2 Cell[K]) bool) Grid[K] {
	g := Grid[K]{
		Width:      len(cells[0]),
		Height:     len(cells),
		connectFun: connectFun,
	}

	cellGrid := make([][]*Cell[K], g.Height)
	for i := range cellGrid {
		cellGrid[i] = make([]*Cell[K], g.Width)
	}

	g.Cells = cells

	return g
}

func (g Grid[K]) Reset(resetFun func(context K) K) {
	for _, row := range g.Cells {
		for _, p := range row {
			p.Context = resetFun(p.Context)
			p.Marked = false
		}
	}
}

func (g Grid[K]) GetCell(i Index) *Cell[K] {
	return g.Cells[i.Y][i.X]
}

func (g Grid[K]) InGrid(i Index) bool {
	return i.X >= 0 && i.X < g.Width && i.Y >= 0 && i.Y < g.Height
}

func (g Grid[K]) PrintBeautifully(cells ...*Cell[K]) {
	c := color.New(color.FgCyan, color.Bold)
	r := color.New(color.FgRed, color.Bold)
	fmt.Println()
	for _, row := range g.Cells {
		for _, p := range row {
			if slices.Contains(cells, p) {
				r.Print(p.Repr)
			} else if (*p).Marked {
				c.Print(p.Repr)
			} else {
				fmt.Printf(fmt.Sprint(p.Repr))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid[K]) Print() {
	fmt.Println()
	for _, row := range g.Cells {
		for _, p := range row {
			fmt.Printf(fmt.Sprint((*p).String()))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid[K]) PrintContext(pFunc func(c Cell[K]) string) {
	fmt.Println()
	for _, row := range g.Cells {
		for _, p := range row {
			fmt.Printf(pFunc(*p))
		}
		fmt.Println()
	}
	fmt.Println()
}

// func (g Grid) PrintMarked() {
// 	fmt.Println()
// 	for _, row := range g.Points {
// 		for _, p := range row {
// 			if p.Marked {
// 				fmt.Printf(fmt.Sprint("#"))
// 			} else {
// 				fmt.Printf(fmt.Sprint(" "))
// 			}
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

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
			Y: i.Y + 1,
		},
	}
}
