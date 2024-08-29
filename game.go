package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type game struct {
	cells [][]bool
}

// New creates a New Game of rows/cols size and inits the game with random life
func (g game) New(rows, cols int) *game {
	n := game{}
	for r := 0; r < rows; r++ {
		row := []bool{}
		for c := 0; c < cols; c++ {
			if rand.Int()%50 == 0 {
				row = append(row, true)
				continue
			}
			row = append(row, false)
		}
		n.cells = append(n.cells, row)
	}
	return &n
}

// Colors returns a slice of tcell.Color with dead cells as black and live cells as white
func (g *game) Colors() [][]tcell.Color {
	colors := [][]tcell.Color{}
	for _, row := range g.cells {
		rc := []tcell.Color{}
		for _, life := range row {
			color := tcell.ColorBlack
			if life {
				color = tcell.ColorWhite
			}
			rc = append(rc, color)
		}
		colors = append(colors, rc)
	}
	return colors
}

// Play moves the game forward one itteration
func (g *game) Play() {
	n := [][]bool{}
	for r, row := range g.cells {
		nRow := []bool{}
		for c, life := range row {
			count := g.countLiving(r, c)
			nLife := life
			if life {
				if count < 2 || count > 3 {
					nLife = false
				}
			} else if count == 3 {
				nLife = true
			}
			nRow = append(nRow, nLife)
		}
		n = append(n, nRow)
	}
	g.cells = n
}

// countLiving will determine how many living cells are around a specific cell. Used during game mechanics
func (g *game) countLiving(r, c int) int {
	count := 0

	pos := [][]int{
		{r - 1, c - 1}, {r - 1, c}, {r - 1, c + 1},
		{r, c - 1}, {r, c + 1},
		{r + 1, c - 1}, {r + 1, c}, {r + 1, c + 1},
	}

	for _, p := range pos {
		if !g.valid(p[0], p[1]) {
			continue
		}

		if g.cells[p[0]][p[1]] {
			count++
		}
	}
	return count
}

// valid is a hler function that will return true if the cell you are attempting to reference is on the board
func (g *game) valid(r, c int) bool {
	return r >= 0 && r < len(g.cells) && c >= 0 && c < len(g.cells[0])
}

// set will set the state of the game. Useful when trying to load in seed files
func (g *game) set(in [][]bool) {
	for r, row := range g.cells {
		for c := range row {
			g.cells[r][c] = false
			if r < len(in) && c < len(in[r]) {
				g.cells[r][c] = in[r][c]
			}
		}
	}
}
