package bs

// searchInsert 35. 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-interview-150
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r + 1
}
