package queue

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum/
func maxSlidingWindow1(nums []int, k int) []int {
	monoQueue := make([]int, 0, len(nums)-k+1)
	if len(nums) < k {
		return []int{}
	}

	for i := 0; i < k; i++ {
		for len(monoQueue) > 0 && nums[i] >= nums[monoQueue[len(monoQueue)-1]] {
			monoQueue = monoQueue[:len(monoQueue)-1]
		}
		monoQueue = append(monoQueue, i)
	}

	ans, ansI := make([]int, len(nums)-k+1), 0
	ans[ansI] = nums[monoQueue[0]]
	ansI++

	for i := k; i < len(nums); i++ {
		for len(monoQueue) > 0 && monoQueue[0] < i-k+1 {
			monoQueue = monoQueue[1:]
		}

		for len(monoQueue) > 0 && nums[i] >= nums[monoQueue[len(monoQueue)-1]] {
			monoQueue = monoQueue[:len(monoQueue)-1]
		}
		monoQueue = append(monoQueue, i)

		ans[ansI] = nums[monoQueue[0]]
		ansI++
	}

	return ans
}
