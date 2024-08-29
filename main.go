package main

import (
	"Conway/board"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	row, col := 13, 40

	board := board.New(row, col)
	app.SetRoot(board.Grid, true)

	g := game{}.New(row, col)

	seed := loadSeed(SD_GOSPER_GLIDER_GUN)
	seed, err := padToCenter(seed, row, col)

	if err != nil {
		log.Panic(err)
	}

	g.set(seed)

	board.Paint([][]tcell.Color{})

	go Loop(app, board, g)

	if err := app.Run(); err != nil {
		panic(err)
	}

}

// Loop: handles triggering of game logic and painting of board. Paced to once every 100ms
func Loop(app *tview.Application, b *board.Board, g *game) {
	for {
		time.Sleep(time.Millisecond * 100)
		g.Play()
		b.Paint(g.Colors())
		app.ForceDraw()
	}
}
