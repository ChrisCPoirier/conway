package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_game_countLiving(t *testing.T) {
	type fields struct {
		cells [][]bool
	}
	type args struct {
		r int
		c int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: `zero neighbors`,
			fields: fields{cells: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			}},
			args: args{r: 1, c: 1},
			want: 0,
		},
		{
			name: `one neighbors`,
			fields: fields{cells: [][]bool{
				{false, false, true},
				{false, true, false},
				{false, false, false},
			}},
			args: args{r: 1, c: 1},
			want: 1,
		},
		{
			name: `all neighbors`,
			fields: fields{cells: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			}},
			args: args{r: 1, c: 1},
			want: 8,
		},
		{
			name: `all neighbors`,
			fields: fields{cells: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			}},
			args: args{r: 1, c: 0},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &game{
				cells: tt.fields.cells,
			}
			got := g.countLiving(tt.args.r, tt.args.c)
			assert.Equal(t, tt.want, got)
		})
	}
}
