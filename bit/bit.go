package bit

// 67. 二进制求和
// https://leetcode.cn/problems/add-binary/description/?envType=study-plan-v2&envId=top-interview-150
func addBinary(a string, b string) string {
	aI, bI := len(a)-1, len(b)-1
	i := max(aI, bI) + 1
	ret := make([]byte, i+1)
	mod, tmp := byte(0), byte(0)
	for aI >= 0 && bI >= 0 {
		tmp = a[aI] + b[bI] - 2*'0' + mod
		ret[i] = tmp%2 + '0'
		mod = tmp / 2
		i--
		aI--
		bI--
	}

	for aI >= 0 {
		tmp = a[aI] - '0' + mod
		ret[i] = tmp%2 + '0'
		mod = tmp / 2
		i--
		aI--
	}
	for bI >= 0 {
		tmp = b[bI] - '0' + mod
		ret[i] = tmp%2 + '0'
		mod = tmp / 2
		i--
		bI--
	}
	if mod > 0 {
		ret[i] = mod + '0'
		return string(ret)
	}
	return string(ret[1:])
}

// 190. 颠倒二进制位
// https://leetcode.cn/problems/reverse-bits/?envType=study-plan-v2&envId=top-interview-150
func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	for i := 0; i < 32; i++ {
		ans = ans*2 + (num & 1)
		num >>= 1
	}
	return ans
}

// 191. 位1的个数
// https://leetcode.cn/problems/number-of-1-bits/description/?envType=study-plan-v2&envId=top-interview-150
func hammingWeight(num int) int {
	ans := 0
	for i := 0; i < 31; i++ {
		ans += num & 1
		num >>= 1
	}
	return ans
}

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/?envType=study-plan-v2&envId=top-interview-150
func singleNumber(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

// 9. 回文数
// https://leetcode.cn/problems/palindrome-number/description/?envType=study-plan-v2&envId=top-interview-150
func isPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}
