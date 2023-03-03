package dp

// uniquePaths
// 62. 不同路径
// https://leetcode.cn/problems/unique-paths/
func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	rows := make([][]int, m)
	for i := 0; i < m; i++ {
		rows[i] = make([]int, n)
		if i == 0 {
			rows[0][0] = 1
		}
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				rows[i][j] += rows[i-1][j]
			}
			if j-1 >= 0 {
				rows[i][j] += rows[i][j-1]
			}
		}
	}

	return rows[m-1][n-1]
}
