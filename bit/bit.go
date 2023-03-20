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

// reverseBits
// 190. 颠倒二进制位
// https://leetcode.cn/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	flagEnd, flagBegin, sum := uint32(1), uint32(1<<31), uint32(0)
	for i := 0; i < 32; i++ {
		if num&flagEnd == flagEnd {
			sum += flagBegin
		}
		flagBegin >>= 1
		flagEnd <<= 1
	}

	return sum
}

// countBits
// 338. 比特位计数
// https://leetcode.cn/problems/counting-bits/
func countBits(n int) []int {
	if n < 1 {
		return nil
	}

	ans := make([]int, n+1)
	ans[0] = 0
	highBit := 0

	for i := 1; i < n+1; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		ans[i] = ans[i-highBit] + 1
	}
	return ans
}
