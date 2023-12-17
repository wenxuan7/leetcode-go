package array

// findMedianSortedArrays
// 4. 寻找两个正序数组的中位数
// https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	prev, curr := 0, 0
	i, j, l := 0, 0, len(nums1)+len(nums2)
	count := 0
	for count <= l/2 {
		prev = curr
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				curr = nums1[i]
				i++
			} else {
				curr = nums2[j]
				j++
			}
		} else if i < len(nums1) {
			curr = nums1[i]
			i++
		} else if j < len(nums2) {
			curr = nums2[j]
			j++
		}
		count++
	}

	if l%2 == 0 {
		return float64(prev+curr) / 2
	} else {
		return float64(curr)
	}
}
