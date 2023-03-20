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

func TestIsPowerOfTwo(t *testing.T) {
	n := []int{
		1,
		16,
		3,
		4,
		5,
	}
	actual := []bool{
		true,
		true,
		false,
		true,
		false,
	}

	for i := range n {
		result := isPowerOfTwo(n[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestReverseBits(t *testing.T) {
	num := []uint32{
		0b00000010100101000001111010011100,
		0b11111111111111111111111111111101,
	}
	actual := []uint32{
		964176192,
		3221225471,
	}

	for i := range num {
		result := reverseBits(num[i])
		assert.Verify(t, i, result, actual[i])
	}
}
