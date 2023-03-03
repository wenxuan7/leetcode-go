package dp

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	m := []int{
		3,
		3,
	}
	n := []int{
		7,
		2,
	}
	actual := []int{
		28,
		3,
	}

	for i := range m {
		result := uniquePaths(m[i], n[i])
		assert.Verify(t, i, result, actual[i])
	}
}
