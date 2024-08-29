package board

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Board struct {
	Grid  *tview.Grid
	cells [][]*tview.Box
}

func New(rows, cols int) *Board {
	b := &Board{}
	newCell := func() *tview.Box {
		color := tcell.ColorBlack
		return tview.NewBox().
			SetBackgroundColor(color).
			SetBorderPadding(0, 0, 0, 0)
	}

	b.Grid = tview.NewGrid().
		SetBorders(true).
		SetSize(rows, cols, 1, 1)

	for r := 0; r < rows; r++ {
		row := []*tview.Box{}
		for c := 0; c < cols; c++ {
			cell := newCell()
			row = append(row, cell)
			b.Grid.AddItem(cell, r, c, 1, 1, 0, 0, false)
		}
		b.cells = append(b.cells, row)
	}

	return b
}

func (g *Board) Paint(colors [][]tcell.Color) *Board {
	for r, row := range colors {
		for c, color := range row {
			g.cells[r][c].SetBackgroundColor(color)
		}
	}
	return g
}
