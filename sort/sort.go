package sort

// sortColors
// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors/
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	p0, p2 := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			if nums[i] == 2 {
				nums[i], nums[p2] = nums[p2], nums[i]
			}
		}

		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
