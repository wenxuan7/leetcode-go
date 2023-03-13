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
