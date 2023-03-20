package bit

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	num := []uint32{
		0b00000000000000000000000000001011,
		0b00000000000000000000000010000000,
		0b11111111111111111111111111111101,
	}
	actual := []int{
		3,
		1,
		31,
	}

	for i := range num {
		result := hammingWeight(num[i])
		assert.Verify(t, i, result, actual[i])
	}
}
