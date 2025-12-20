package alg

import (
	"math/rand/v2"
	"slices"
)

// MergeSort 合并排序
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) >> 1
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	res := make([]int, 0, len(arr))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res
}

// QuickSort 快速排序
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	moveRandomToFirst(arr)
	pivot, l := 0, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[pivot] {
			arr[l], arr[i] = arr[i], arr[l]
			l++
		}
	}
	arr[pivot], arr[l-1] = arr[l-1], arr[pivot]

	QuickSort(arr[:l])
	QuickSort(arr[l:])
}

func moveRandomToFirst(arr []int) {
	n := rand.IntN(len(arr))
	arr[0], arr[n] = arr[n], arr[0]
}

// CountingSort 计数排序
func CountingSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mn, mx := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		mn = min(mn, arr[i])
		mx = max(mx, arr[i])
	}

	L := mx - mn + 1
	cnt := make([]int, L)
	for i := 0; i < len(arr); i++ {
		cnt[arr[i]-mn]++
	}
	for i := 1; i < L; i++ {
		cnt[i] += cnt[i-1]
	}

	res := make([]int, len(arr))
	num, sub, idx := 0, 0, 0
	for r := len(arr) - 1; r >= 0; r-- {
		num = arr[r]
		sub = num - mn
		idx = cnt[sub] - 1

		cnt[sub] = idx
		res[idx] = num
	}
	return res
}

// BucketSort 桶排序
func BucketSort(arr []int, bucketSize int) []int {
	if len(arr) < 2 {
		return arr
	}

	mn, mx := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		mn = min(mn, arr[i])
		mx = max(mx, arr[i])
	}

	// (mx-mn+1-1)/bucketSize + 1
	bucketCnt := (mx-mn)/bucketSize + 1
	buckets := make([][]int, bucketCnt)
	idx := 0
	for _, num := range arr {
		idx = (num - mn) / bucketSize
		buckets[idx] = append(buckets[idx], num)
	}

	res := make([]int, 0, len(arr))
	var sorted []int
	for _, nums := range buckets {
		sorted = CountingSort(nums)
		res = append(res, sorted...)
	}
	return res
}

// RadixSort 基数排序
func RadixSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mx := slices.Max(arr)
	res := arr

	for exp := 1; mx/exp > 0; exp *= 10 {
		res = radixSort(res, exp)
	}
	return res
}

func radixSort(arr []int, exp int) []int {
	cnt := make([]int, 10)
	for _, num := range arr {
		cnt[(num/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		cnt[i] += cnt[i-1]
	}

	res := make([]int, len(arr))
	idx, num, cntIdx := 0, 0, 0
	for r := len(arr) - 1; r >= 0; r-- {
		num = arr[r]
		cntIdx = (num / exp) % 10
		idx = cnt[cntIdx] - 1

		cnt[cntIdx] = idx
		res[idx] = num
	}
	return res
}
