package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const (
	SD_GOSPER_GLIDER_GUN = `seeds/gosperGliderGun.txt`
	SD_STAR              = `seeds/star.txt`
)

// loadSeed will take in the location of a seed file and turn it into a matrix of bools representing life or death
func loadSeed(filename string) [][]bool {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	seed := [][]bool{}
	for _, line := range strings.Split(string(bytes), "\n") {
		if len(line) > 0 && line[0] == '!' {
			continue
		}
		row := []bool{}
		for _, r := range line {
			if r == 'O' {
				row = append(row, true)
				continue
			}
			row = append(row, false)
		}
		seed = append(seed, row)
	}

	return seed
}

type ErrPadOutOfBounds struct {
	In, Board   int
	RowOrColumn string
}

func (e ErrPadOutOfBounds) Error() string {
	return fmt.Sprintf("cannot pad input %s of %d with board %s %d. Input longer than board", e.RowOrColumn, e.In, e.RowOrColumn, e.Board)
}

// in is the slice of bool, row and column are the size we want to adjust the in slice to
func padToCenter(in [][]bool, row, column int) ([][]bool, error) {
	if row < len(in) {
		return [][]bool{}, ErrPadOutOfBounds{In: len(in), Board: row, RowOrColumn: `row`}
	}

	if column < len(in[0]) {
		return [][]bool{}, ErrPadOutOfBounds{In: len(in[0]), Board: column, RowOrColumn: `row`}
	}

	row = centerIndex(row)
	column = centerIndex(column)

	inRow := centerIndex(len(in))
	inColumn := centerIndex(longestColumn(in))

	rowToPad := row - inRow
	colToPad := column - inColumn

	if rowToPad > 0 {
		in = append(make([][]bool, rowToPad), in...)
	}

	if colToPad <= 0 {
		return in, nil
	}

	for i := range in {
		in[i] = append(make([]bool, colToPad), in[i]...)
	}

	return in, nil
}

func centerIndex(i int) int {
	c := math.Ceil(float64(i)/2) - 1
	return int(c)
}

func longestColumn(in [][]bool) int {
	longest := 0
	for _, row := range in {
		if len(row) > longest {
			longest = len(row)
		}
	}
	return longest
}
