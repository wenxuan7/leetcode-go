package str

import (
	"bytes"
	"math"
	"strings"
)

// toLowerCase
// 709. 转换成小写字母
// https://leetcode.cn/problems/to-lower-case/
func toLowerCase(s string) string {
	if s == "" {
		return s
	}

	chs := make([]byte, len(s))
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			chs[i] = s[i] - 'A' + 'a'
			continue
		}
		chs[i] = s[i]
	}
	return string(chs)
}

// lengthOfLastWord
// 58. 最后一个单词的长度
// https://leetcode.cn/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	i, ans := len(s)-1, 0
	for i > -1 {
		if s[i] != ' ' {
			break
		}
		i--
	}

	for i > -1 {
		if s[i] != ' ' {
			ans++
		} else {
			break
		}
		i--
	}
	return ans
}

// numJewelsInStones
// 771. 宝石与石头
// https://leetcode.cn/problems/jewels-and-stones/
func numJewelsInStones(jewels string, stones string) int {
	lower, upper := make([]bool, 26), make([]bool, 26)
	for i := range jewels {
		temp := jewels[i]
		if temp >= 'A' && temp <= 'Z' {
			upper[temp-'A'] = true
			continue
		}
		lower[temp-'a'] = true
	}

	ans := 0
	for i := range stones {
		temp := stones[i]
		if temp >= 'A' && temp <= 'Z' {
			if upper[temp-'A'] {
				ans++
			}
			continue
		}

		if lower[temp-'a'] {
			ans++
		}
	}
	return ans
}

// 387. 字符串中的第一个唯一字符
// https://leetcode.cn/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	if 0 == len(s) {
		return -1
	}

	set := make([]int, 26)
	for i := range set {
		set[i] = len(s)
	}
	for i, ch := range s {
		k := ch - 'a'
		if set[k] == len(s) {
			set[k] = i
		} else {
			set[k] = len(s) + 1
		}
	}

	ans := len(s)
	for _, num := range set {
		if ans > num {
			ans = num
		}
	}
	if ans < len(s) {
		return ans
	}
	return -1
}

// myAtoi
// 8. 字符串转换整数 (atoi)
// https://leetcode.cn/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)

	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}

	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}

	for ; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		result = result*10 + int(s[i]-'0')
		if sign*result < math.MinInt32 {
			return math.MinInt32
		}
		if sign*result > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return sign * result
}

// maskPII
// 831. 隐藏个人信息
// https://leetcode.cn/problems/masking-personal-information/
func maskPII(s string) string {
	if len(s) == 0 {
		return ""
	}

	mailFlagI := strings.IndexByte(s, '@')
	if mailFlagI >= 0 {
		// 名字 和 域名 部分的大写英文字母应当转换成小写英文字母。
		// 名字 中间的字母（即，除第一个和最后一个字母外）必须用 5 个 "*****" 替换。
		bs := make([]byte, 7+len(s)-mailFlagI)
		for i, j := mailFlagI, 7; i < len(s); i, j = i+1, j+1 {
			b := s[i]
			if b >= 'A' && b <= 'Z' {
				bs[j] = b - 'A' + 'a'
			} else {
				bs[j] = b
			}
		}

		for i := 1; i < 6; i++ {
			bs[i] = '*'
		}

		s0 := s[0]
		if s0 >= 'A' && s0 <= 'Z' {
			bs[0] = s0 - 'A' + 'a'
		} else {
			bs[0] = s0
		}

		s1 := s[mailFlagI-1]
		if s1 >= 'A' && s1 <= 'Z' {
			bs[6] = s1 - 'A' + 'a'
		} else {
			bs[6] = s1
		}

		return string(bs)
	}

	// ***-***-XXXX 如果国家代码为 0 位数字
	// +*-***-***-XXXX 如果国家代码为 1 位数字
	// +**-***-***-XXXX 如果国家代码为 2 位数字
	// +***-***-***-XXXX 如果国家代码为 3 位数字
	numCount := 0
	last4, j := []byte{0, 0, 0, 0}, 3
	for i := len(s) - 1; i > -1; i-- {
		b := s[i]
		if b >= '0' && b <= '9' {
			numCount++
			if j > -1 {
				last4[j] = b
				j--
			}
		}
	}

	var ans []byte
	switch numCount {
	case 10:
		ans = []byte{'*', '*', '*', '-', '*', '*', '*', '-', 'X', 'X', 'X', 'X'}
	case 11:
		ans = []byte{'+', '*', '-', '*', '*', '*', '-', '*', '*', '*', '-', 'X', 'X', 'X', 'X'}
	case 12:
		ans = []byte{'+', '*', '*', '-', '*', '*', '*', '-', '*', '*', '*', '-', 'X', 'X', 'X', 'X'}
	case 13:
		ans = []byte{'+', '*', '*', '*', '-', '*', '*', '*', '-', '*', '*', '*', '-', 'X', 'X', 'X', 'X'}
	default:
		ans = []byte{'*', '*', '*', '-', '*', '*', '*', '-', 'X', 'X', 'X', 'X'}
	}

	copy(ans[len(ans)-4:], last4)
	return string(ans)
}

// isValid
// 1003. 检查替换后的词是否有效
// https://leetcode.cn/problems/check-if-word-is-valid-after-substitutions/
func isValid(s string) bool {
	if len(s)%3 != 0 {
		return false
	}

	stk := make([]byte, len(s))
	top := -1
	for i := range s {
		b := s[i]
		top++
		stk[top] = b

		if top >= 2 &&
			stk[top] == 'c' &&
			stk[top-1] == 'b' &&
			stk[top-2] == 'a' {
			top -= 3
		}
	}

	return top == -1
}

// gcdOfStrings
// 1071. 字符串的最大公因子
// https://leetcode.cn/problems/greatest-common-divisor-of-strings
// TODO 其他优秀的解法没有思考过
func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		return gcdOfStrings(str2, str1)
	}

	bs1, bs2 := []byte(str1), []byte(str2)
	for i := len(bs1); i >= 1; i-- {
		if len(str1)%i != 0 ||
			len(str2)%i != 0 {
			continue
		}

		common := bs1[0:i]

		flag := true
		for j := 0; j < len(bs1)/i; j++ {
			temp := bs1[j*i : (j+1)*i]

			if !bytes.Equal(common, temp) {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		for k := 0; k < len(bs2)/i; k++ {
			temp := bs2[k*i : (k+1)*i]

			if !bytes.Equal(common, temp) {
				flag = false
				break
			}
		}
		if flag {
			return string(common)
		}
	}
	return ""
}

// replaceSpace
// 剑指 Offer 05. 替换空格
// https://leetcode.cn/problems/ti-huan-kong-ge-lcof
func replaceSpace(s string) string {
	if s == "" {
		return ""
	}

	bs := make([]byte, 0, len(s)*3)
	for i := range s {
		b := s[i]
		if b != ' ' {
			bs = append(bs, b)
		} else {
			bs = append(bs, "%20"...)
		}
	}

	return string(bs)
}

// replaceWords
// 648. 单词替换
// https://leetcode.cn/problems/replace-words/description/
func replaceWords(dictionary []string, sentence string) string {
	qSortStr(dictionary, 0, len(dictionary)-1)
	split := strings.Split(sentence, " ")
	ans := make([]string, 0, len(split))
	for _, s := range split {
		flag := false
		for _, dic := range dictionary {
			if len(dic) > len(s) {
				break
			}

			if s[:len(dic)] == dic {
				ans = append(ans, dic)
				flag = true
				break
			}
		}
		if !flag {
			ans = append(ans, s)
		}
	}

	return strings.Join(ans, " ")
}

func qSortStr(s []string, l, r int) {
	if l >= r {
		return
	}

	flag, index := l, l+1
	for i := l + 1; i <= r; i++ {
		if len(s[i]) < len(s[flag]) {
			s[index], s[i] = s[i], s[index]
			index++
		}
	}

	s[flag], s[index-1] = s[index-1], s[flag]
	flag = index - 1
	qSortStr(s, l, flag-1)
	qSortStr(s, flag+1, r)
}
