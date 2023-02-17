package recursive

import (
	"leetcode-go/assert"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	data := [][]int{
		{1, 2, 3, 1, 1},
	}
	actual := []int{
		1,
	}

	for i := range data {
		result := majorityElement(data[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestLetterCombinations(t *testing.T) {
	data := []string{
		"23",
		"",
	}
	actual := []map[string]bool{
		{
			"ad": true,
			"ae": true,
			"af": true,
			"bd": true,
			"be": true,
			"bf": true,
			"cd": true,
			"ce": true,
			"cf": true,
		},
		{},
	}

	for i := range data {
		result := letterCombinations(data[i])
		if len(result) != len(actual[i]) {
			t.Fatalf("结果与实际不相符, resultLen: %d, actualLen: %d", len(result), len(actual[i]))
		}

		for k := range result {
			b, ok := actual[i][result[k]]
			if !ok {
				t.Fatalf("结果与实际不相符, result: %s", result[k])
			}
			if !b {
				t.Fatalf("结果有重复, result: %s", result[k])
			}

			actual[i][result[k]] = false
		}
	}
}

func TestSolveNQueens(t *testing.T) {
	data := []int{
		4,
		5,
	}

	for i := range data {
		result := solveNQueens(data[i])
		t.Logf("%v\n", result)
	}
}
