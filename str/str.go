package str

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
