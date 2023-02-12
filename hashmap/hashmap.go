package hashmap

// isAnagram
// 242. 有效的字母异位词
// https://leetcode.cn/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	m := [26]int{}
	for _, b := range []byte(s) {
		m[b-'a']++
	}

	for _, b := range []byte(t) {
		m[b-'a']--
		if m[b-'a'] < 0 {
			return false
		}
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}
