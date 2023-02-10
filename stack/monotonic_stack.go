package stack

// largestRectangleArea
// 84. 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	monoStack := make([]int, len(heights))
	top := -1
	l, r := make([]int, len(heights)), make([]int, len(heights))
	// 初始化边界
	for i := range heights {
		r[i] = len(heights)
	}

	for i := range heights {
		for top > -1 && heights[monoStack[top]] >= heights[i] {
			r[monoStack[top]] = i
			top--
		}

		if top < 0 {
			l[i] = -1
		} else {
			l[i] = monoStack[top]
		}

		top++
		monoStack[top] = i
	}

	max := 0
	for i := range l {
		area := (r[i] - l[i] - 1) * heights[i]
		if area > max {
			max = area
		}
	}

	return max
}
