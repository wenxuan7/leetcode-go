package recursive

import (
	"github.com/leetcode-go/assert"
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

	for i := range data {
		result := numIslands(data[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestUpdateBoard(t *testing.T) {
	data := [][][]byte{
		{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
		},
	}
	click := [][]int{
		{3, 0},
	}
	actual := [][][]byte{
		{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
	}

	for i := range data {
		result := updateBoard(data[i], click[i])
		assert.Verify2Arr(t, i, result, actual[i])
	}
}
