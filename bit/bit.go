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
