package math

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestSmallestEvenMultiple(t *testing.T) {
	num := []int{
		5,
		6,
		7,
	}
	actual := []int{
		10,
		6,
		14,
	}

	for i := range num {
		result := smallestEvenMultiple(num[i])
		assert.Verify(t, i, result, actual[i])
	}
}
