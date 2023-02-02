package array

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}

	l, r, max := 0, len(height)-1, 0
	for l < r {
		wid, area := r-l, 0
		if height[r] > height[l] {
			area = height[l] * wid
			oldL := l
			for l < r && height[l] <= height[oldL] {
				l++
			}
		} else {
			area = height[r] * wid
			oldR := r
			for l < r && height[r] <= height[oldR] {
				r--
			}
		}

		if area > max {
			max = area
		}
	}

	return max
}
