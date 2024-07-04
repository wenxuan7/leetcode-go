package array

import (
	"container/heap"
	"math"
)

// 4. 寻找两个正序数组的中位数
// findMedianSortedArrays https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	l, r := 0, len(nums1)
	sumL := len(nums1) + len(nums2)
	mid := (sumL + 1) >> 1
	for l <= r {
		mid1 := (l + r + 1) >> 1
		mid2 := mid - mid1
		leftMax1, leftMax2 := math.MinInt, math.MinInt
		rightMin1, rightMin2 := math.MaxInt, math.MaxInt
		if mid1 != 0 {
			leftMax1 = nums1[mid1-1]
		}
		if mid1 != len(nums1) {
			rightMin1 = nums1[mid1]
		}

		if mid2 != 0 {
			leftMax2 = nums2[mid2-1]
		}
		if mid2 != len(nums2) {
			rightMin2 = nums2[mid2]
		}

		if leftMax1 <= rightMin2 && leftMax2 <= rightMin1 {
			if sumL%2 == 0 {
				return float64(max(leftMax1, leftMax2)+min(rightMin1, rightMin2)) / 2.0
			}
			return float64(max(leftMax1, leftMax2))
		} else if leftMax1 > rightMin2 {
			r = mid1 - 1
		} else {
			l = mid1 + 1
		}
	}
	return 0.0
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func NewItem(value, priority int) *Item {
	return &Item{value, priority}
}

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	p := *pq
	return p[i].priority > p[j].priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	p := *pq
	p[i], p[j] = p[j], p[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	pq := make(PriorityQueue, len(m))
	i := 0
	for value, priority := range m {
		pq[i] = NewItem(value, priority)
		i++
	}
	heap.Init(&pq)

	ans := make([]int, 0, k)
	for j := 0; j < k; j++ {
		item := heap.Pop(&pq).(*Item)
		ans = append(ans, item.value)
	}
	return ans
}
