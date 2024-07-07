package graph

// numIslands 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands/description/?envType=study-plan-v2&envId=top-interview-150
func numIslands(grid [][]byte) int {
	ret := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '0' {
				continue
			}
			reset(grid, r, c)
			ret++
		}
	}
	return ret
}

func reset(grid [][]byte, r, c int) {
	if r < 0 ||
		c < 0 ||
		r >= len(grid) ||
		c >= len(grid[0]) {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	reset(grid, r-1, c)
	reset(grid, r+1, c)
	reset(grid, r, c-1)
	reset(grid, r, c+1)
}
