package unionset

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	isConnected := [][][]int{
		{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
		{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
	}
	actual := []int{
		2,
		3,
	}

	for i := range isConnected {
		result := findCircleNum(isConnected[i])
		assert.Verify(t, i, result, actual[i])
	}
}

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

	for i := range data {
		result := numIslands(data[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestSolve(t *testing.T) {
	data := [][][]byte{
		{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		},
		{
			{'X'},
		},
	}
	actual := [][][]byte{
		{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		},
		{
			{'X'},
		},
	}

	for i := range data {
		solve(data[i])
		assert.Verify2Arr(t, i, data[i], actual[i])
	}
}
