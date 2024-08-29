package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_padToCenter(t *testing.T) {
	type args struct {
		in     [][]bool
		row    int
		column int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]bool
		wantErr error
	}{
		{
			name: `board is smaller than input`,
			args: args{
				in: [][]bool{
					{true, true},
					{true, true},
				},
				row:    1,
				column: 1,
			},
			want:    [][]bool{},
			wantErr: ErrPadOutOfBounds{In: 2, Board: 1, RowOrColumn: `row`},
		},
		{
			name: `board is equal to input`,
			args: args{
				in: [][]bool{
					{true, true},
					{true, true},
				},
				row:    2,
				column: 2,
			},
			want: [][]bool{
				{true, true},
				{true, true},
			},
		},
		{
			name: `even board, even input`,
			args: args{
				in: [][]bool{
					{true, true},
					{true, true},
				},
				row:    8,
				column: 8,
			},
			want: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
				{false, false, false, true, true},
				{false, false, false, true, true},
			},
		},
		{
			name: `even board, un-even input`,
			args: args{
				in: [][]bool{
					{true, true, true},
					{true, true, true},
					{true, true, true},
				},
				row:    8,
				column: 8,
			},
			want: [][]bool{
				{false, false},
				{false, false},
				{false, false, true, true, true},
				{false, false, true, true, true},
				{false, false, true, true, true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := padToCenter(tt.args.in, tt.args.row, tt.args.column)

			if tt.wantErr != nil || err != nil {
				assert.EqualError(t, err, tt.wantErr.Error(), tt.name)
			}

			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}
