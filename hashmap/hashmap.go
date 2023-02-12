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

// groupAnagrams
// 49. 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
