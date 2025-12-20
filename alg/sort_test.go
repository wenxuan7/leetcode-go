package alg

import "testing"

func TestSort(t *testing.T) {
	numsArr := [][]int{
		{5, 4, 3, 2, 1},
		{10, 11, 8, 12, 7, 6, 5, 4, 1, 3, 2},
		{33, 43, 1, 10, 5, 2, 1},
	}

	for _, nums := range numsArr {
		merge := MergeSort(nums)
		t.Log("merge:", merge)

		count := CountingSort(nums)
		t.Log("count:", count)

		bucket := BucketSort(nums, 5)
		t.Log("bucket:", bucket)

		radix := RadixSort(nums)
		t.Log("radix:", radix)

		QuickSort(nums)
		t.Log("quick:", nums)
	}
}
