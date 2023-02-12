package stack

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	var (
		data = [][]int{
			{2, 1, 5, 6, 2, 3},
			{1, 2, 2, 2, 1, 3, 4},
		}
		actual = []int{
			10,
			7,
		}
	)

	for i := range data {
		result := largestRectangleArea(data[i])
		if result != actual[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %d, actual: %d",
				i, result, actual[i])
		}
	}
}

func TestTrap(t *testing.T) {
	data := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
	}
	actual := []int{
		6,
	}

	for i := range data {
		result := trap(data[i])
		if result != actual[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %d, actual: %d",
				i, result, actual[i])
		}
	}
}
