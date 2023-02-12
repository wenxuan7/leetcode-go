package queue

import "testing"

func TestMaxSlidingWindow1(t *testing.T) {
	data := [][]int{
		{1},
		{1, 3, -1, -3, 5, 3, 6, 7},
	}
	k := []int{
		1,
		3,
	}
	actual := [][]int{
		{1},
		{3, 3, 5, 5, 6, 7},
	}

	for i := range data {
		result := maxSlidingWindow1(data[i], k[i])
		verify(t, i, result, actual[i])
	}
}

func verify(t *testing.T, caseIndex int, result []int, actual []int) {
	if result == nil && actual == nil {
		return
	}

	if len(result) != len(actual) {
		t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, resultLen: %d, actual: %v, actualLen: %d",
			caseIndex, result, len(result), actual, len(actual))
	}
	for i := range actual {
		if actual[i] != result[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v, i: %d",
				caseIndex, result, actual, i)
		}
	}
}
