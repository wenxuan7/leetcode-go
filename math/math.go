package math

// 66. 加一
// https://leetcode.cn/problems/plus-one/description/?envType=study-plan-v2&envId=top-interview-150
func plusOne(digits []int) []int {
	mod := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += mod
		mod = digits[i] / 10
		digits[i] %= 10
	}
	if mod != 0 {
		ret := make([]int, 0, len(digits)+1)
		ret = append(ret, mod)
		ret = append(ret, digits...)
		return ret
	}
	return digits
}
