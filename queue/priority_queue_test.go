package queue

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
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
