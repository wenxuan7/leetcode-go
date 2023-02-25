package greedy

import (
	"leetcode-go/assert"
	"testing"
)

func TestLemonadeChange(t *testing.T) {
	data := [][]int{
		{5, 5, 5, 10, 20},
		{5, 5, 10, 10, 20},
	}
	actual := []bool{
		true,
		false,
	}

	for i := range data {
		result := lemonadeChange(data[i])
		assert.Verify(t, i, result, actual[i])
	}
}
