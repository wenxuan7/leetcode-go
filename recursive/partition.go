package recursive

// majorityElement
// 169. 多数元素
// https://leetcode.cn/problems/majority-element/description/
func majorityElement(nums []int) int {
	var partition func(nums []int, left, right int) int
	partition = func(nums []int, left, right int) int {
		if left == right {
			return left
		}

		mid := ((right - left) >> 1) + left
		leftMajor := partition(nums, left, mid)
		rightMajor := partition(nums, mid+1, right)

		if leftMajor == rightMajor {
			return leftMajor
		}

		leftMajorCount, rightMajorCount := 0, 0
		for i := left; i <= right; i++ {
			if nums[i] == nums[leftMajor] {
				leftMajorCount++
			}

			if nums[i] == nums[rightMajor] {
				rightMajorCount++
			}
		}

		if leftMajorCount > rightMajorCount {
			return leftMajor
		} else {
			return rightMajor
		}
	}

	ans := partition(nums, 0, len(nums)-1)
	return nums[ans]
}

// letterCombinations
// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	numToLetter := map[int][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}

	bs := []byte(digits)
	var recursive func(start int, temp []byte)
	count := 1
	for _, b := range bs {
		count *= len(numToLetter[int(b-'0')])
	}
	ans := make([]string, 0, count)
	recursive = func(start int, temp []byte) {
		if start == len(bs) {
			dst := make([]byte, len(temp))
			copy(dst, temp)
			ans = append(ans, string(dst))
			return
		}

		b := bs[start]
		for _, b2 := range numToLetter[int(b-'0')] {
			temp = append(temp, b2)
			recursive(start+1, temp)
			temp = temp[:len(temp)-1]
		}
	}

	recursive(0, make([]byte, 0, len(bs)))
	return ans
}

// solveNQueens
// 51. N 皇后
// https://leetcode.cn/problems/n-queens/
func solveNQueens(n int) [][]string {
	if n < 1 {
		return [][]string{}
	}

	ans := make([][]string, 0)
	rows := make([]int, 0, n)
	colHas := make([]bool, n)
	diagonal1, diagonal2 := map[int]bool{}, map[int]bool{}
	var recursive func(row int)
	recursive = func(row int) {
		if row == n {
			ans = append(ans, generateQString(rows))
			return
		}

		for i := 0; i < n; i++ {
			if colHas[i] {
				continue
			}
			if diagonal1[row+i] {
				continue
			}
			if diagonal2[row-i] {
				continue
			}

			colHas[i] = true
			diagonal1[row+i] = true
			diagonal2[row-i] = true
			rows = append(rows, i)
			recursive(row + 1)
			rows = rows[:len(rows)-1]
			diagonal2[row-i] = false
			diagonal1[row+i] = false
			colHas[i] = false
		}
	}

	recursive(0)

	return ans
}

func generateQString(rows []int) []string {
	if len(rows) == 0 {
		return []string{}
	}

	ans := make([]string, 0, len(rows))
	for _, row := range rows {
		bs := make([]byte, len(rows))
		for i := 0; i < len(rows); i++ {
			if i == row {
				bs[i] = 'Q'
				continue
			}
			bs[i] = '.'
		}
		ans = append(ans, string(bs))
	}

	return ans
}
