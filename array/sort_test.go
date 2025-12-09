package array

import "testing"

func TestMergeSort(t *testing.T) {
	data := [][]int{
		{10, 9, 8, 7, 6, 6, 5, 1, 3, 2},
		{9, 9, 10, 10, 3, 1},
		{7, 8, 9},
	}
	for _, nums := range data {
		t.Log(mergeSort(nums))
	}
}

func TestQuickSort(t *testing.T) {
	data := [][]int{
		{10, 9, 8, 7, 6, 6, 5, 1, 3, 2},
		{9, 9, 10, 10, 3, 1},
		{7, 8, 9},
	}
	for _, nums := range data {
		quickSort(nums)
		t.Log(nums)
	}
}

func TestCountingSort(t *testing.T) {
	data := [][]int{
		{10, 9, 8, 7, 6, 6, 5, 1, 3, 2},
		{9, 9, 10, 10, 3, 1},
		{7, 8, 9},
	}
	for _, nums := range data {
		t.Log(countingSort(nums))
	}
}

func TestBucketSort(t *testing.T) {
	data := [][]int{
		{10, 9, 8, 7, 6, 6, 5, 1, 3, 2},
		{9, 9, 10, 10, 3, 1},
		{7, 8, 9},
	}
	for _, nums := range data {
		t.Log(bucketSort(nums, 3))
	}
}
