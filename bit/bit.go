package bit

// hammingWeight
// 191. 位1的个数
// https://leetcode.cn/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	count := 0
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}
