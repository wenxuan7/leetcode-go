package graph

import (
	"github.com/wenxuan7/leetcode/require"
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

func TestMinMutation(t *testing.T) {
	dataStart := []string{
		"AACCGGTT",
	}
	dataEnd := []string{
		"AAACGGTA",
	}
	dataBank := [][]string{
		{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
	}
	actual := []int{
		2,
	}

	for i := range dataStart {
		ret := minMutation(dataStart[i], dataEnd[i], dataBank[i])
		require.EqualN(t, i, ret, actual[i])
	}
}

func TestLadderLength(t *testing.T) {
	dataStart := []string{
		"hit",
		"hbo",
		"hit",
	}
	dataEnd := []string{
		"cog",
		"qbx",
		"cog",
	}
	dataWords := [][]string{
		{"hot", "dot", "dog", "lot", "log", "cog"},
		{"abo", "hco", "hbw", "ado", "abq", "hcd", "hcj", "hww", "qbq", "qby", "qbz", "qbx", "qbw"},
		{"hot", "dot", "dog", "lot", "log"},
	}
	actual := []int{
		5,
		4,
		0,
	}

	for i := range dataStart {
		ret := ladderLength(dataStart[i], dataEnd[i], dataWords[i])
		require.EqualN(t, i, ret, actual[i])
	}
}
