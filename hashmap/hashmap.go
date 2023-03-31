package hashmap

import "github.com/leetcode-go/tool"

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
func groupAnagrams(ss []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range ss {
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

// twoSum
// 1. 两数之和
// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	numToI := make(map[int]int, len(nums))
	for i, v := range nums {
		numsI, ok := numToI[target-v]
		if ok {
			return []int{numsI, i}
		}

		numToI[v] = i
	}

	return []int{}
}

// findSubarrays
// 2395. 和相等的子数组
// https://leetcode.cn/problems/find-subarrays-with-equal-sum/
func findSubarrays(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		k := nums[i] + nums[i+1]
		if m[k] {
			return true
		}
		m[k] = true
	}

	return false
}

// arithmeticTriplets
// 2367. 算术三元组的数目
// https://leetcode.cn/problems/number-of-arithmetic-triplets/
func arithmeticTriplets(nums []int, diff int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := 0
	for i, j, k := 0, 1, 2; i < len(nums)-2; i++ {
		j = tool.Max(i+1, j)
		for ; j < len(nums)-1 && nums[j]-nums[i] < diff; j++ {
		}
		if j >= len(nums)-1 || nums[j]-nums[i] > diff {
			continue
		}

		k = tool.Max(j+1, k)
		for ; k < len(nums) && nums[k]-nums[j] < diff; k++ {
		}
		if k < len(nums) && nums[k]-nums[j] == diff {
			ans++
		}
	}

	return ans
}
