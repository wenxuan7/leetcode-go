package array

import "math/rand/v2"

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
