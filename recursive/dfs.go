package recursive

// numIslands
// 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	set := make([][]bool, row)
	for i := 0; i < row; i++ {
		set[i] = make([]bool, col)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= row {
			return
		}
		if j < 0 || j >= col {
			return
		}

		if grid[i][j] == '1' && !set[i][j] {
			grid[i][j] = '0'

			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		} else {
			return
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				set[i][j] = true
				dfs(i+1, j)
				dfs(i-1, j)
				dfs(i, j+1)
				dfs(i, j-1)
			}
		}
	}

	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				ans++
			}
		}
	}

	return ans
}
