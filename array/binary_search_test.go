package array

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	x := []int{
		4,
		8,
	}
	actual := []int{
		2,
		2,
	}

	for i := range x {
		result := mySqrt(x[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestIsPerfectSquare(t *testing.T) {
	num := []int{
		16,
		14,
		9,
	}
	actual := []bool{
		true,
		false,
		true,
	}

	for i := range num {
		result := isPerfectSquare(num[i])
		assert.Verify(t, i, result, actual[i])
	}
}
