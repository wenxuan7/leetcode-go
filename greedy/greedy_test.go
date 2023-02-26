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

func TestMaxProfit(t *testing.T) {
	data := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
	}
	actual := []int{
		7,
		4,
	}

	for i := range data {
		result := maxProfit(data[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestFindContentChildren(t *testing.T) {
	g := [][]int{
		{1, 2, 3},
		{1, 2},
	}
	s := [][]int{
		{1, 1},
		{1, 2, 3},
	}
	actual := []int{
		1,
		2,
	}

	for i := range g {
		result := findContentChildren(g[i], s[i])
		assert.Verify(t, i, result, actual[i])
	}
}
