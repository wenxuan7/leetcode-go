package array

import (
	"math/rand/v2"
	"slices"
)

// mergeSort 归并排序
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) >> 1
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	nArr := make([]int, 0, len(arr))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nArr = append(nArr, left[i])
			i++
		} else {
			nArr = append(nArr, right[j])
			j++
		}
	}
	nArr = append(nArr, left[i:]...)
	nArr = append(nArr, right[j:]...)
	return nArr
}

// quickSort 快速排序
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// 随机pivot处理 并且将值放置0下标处
	moveRandomToFirst(arr)
	pivot, slow := 0, 1
	for fast := slow; fast < len(arr); fast++ {
		if arr[fast] < arr[pivot] {
			arr[slow], arr[fast] = arr[fast], arr[slow]
			slow++
		}
	}

	arr[slow-1], arr[pivot] = arr[pivot], arr[slow-1]
	quickSort(arr[:slow-1])
	quickSort(arr[slow:])
}

func moveRandomToFirst(arr []int) {
	n := rand.IntN(len(arr))
	arr[0], arr[n] = arr[n], arr[0]
}

// countingSort 计数排序
func countingSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	mn, mx := slices.Min(arr), slices.Max(arr)
	cnt := mx - mn + 1
	cntArr := make([]int, cnt)

	for _, v := range arr {
		cntArr[v-mn] += 1
	}

	// 前缀和处理 确定位置 保证重复元素在原数组的相对顺序
	for i := 1; i < cnt; i++ {
		cntArr[i] += cntArr[i-1]
	}
	res := make([]int, n)
	cur, idx := 0, 0
	for i := len(arr) - 1; i >= 0; i-- {
		cur = arr[i]
		idx = cur - mn
		cntArr[idx]--
		res[cntArr[idx]] = cur
	}

	return res
}

// bucketSort 桶排序
func bucketSort(arr []int, bucketSize int) []int {
	n := len(arr)
	mn, mx := slices.Min(arr), slices.Max(arr)
	// 这里计算桶数量
	// 原来应该是 (mx-mn+1-1)/bucketSize + 1
	bucketCnt := (mx-mn)/bucketSize + 1

	buckets := make([][]int, bucketCnt)
	for i := 0; i < bucketCnt; i++ {
		buckets[i] = make([]int, 0)
	}

	idx := 0
	for _, v := range arr {
		idx = (v - mn) / bucketSize
		buckets[idx] = append(buckets[idx], v)
	}
	for _, bucket := range buckets {
		quickSort(bucket)
	}
	res := make([]int, 0, n)
	for _, bucket := range buckets {
		res = append(res, bucket...)
	}
	return res
}
