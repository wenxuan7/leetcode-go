package graph

import (
	"github.com/leetcode/require"
	"testing"
)

func TestNumIslands(t *testing.T) {
	data := [][][]byte{
		{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		},
		{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
	}
	actual := []int{
		1,
		3,
	}

	for i, grid := range data {
		ret := numIslands(grid)
		require.EqualN(t, i, ret, actual[i])
	}
}

func TestSolve(t *testing.T) {
	data := [][][]byte{
		{
			{'O', 'O'},
			{'O', 'O'},
		},
		{
			{'X', 'O', 'X'},
			{'X', 'O', 'X'},
			{'X', 'O', 'X'},
		},
	}
	actual := [][][]byte{
		{
			{'O', 'O'},
			{'O', 'O'},
		},
		{
			{'X', 'O', 'X'},
			{'X', 'O', 'X'},
			{'X', 'O', 'X'},
		},
	}
	for i, board := range data {
		solve(board)
		require.Equal2D(t, i, board, actual[i])
	}
}
