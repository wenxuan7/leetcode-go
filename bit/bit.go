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

// isPowerOfTwo
// 231. 2 的幂
// https://leetcode.cn/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	return (n & (n - 1)) == 0
}
