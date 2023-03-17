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

// updateBoard
// 529. 扫雷游戏
// https://leetcode.cn/problems/minesweeper/
func updateBoard(board [][]byte, click []int) [][]byte {
	row, col := len(board), len(board[0])
	i, j := click[0], click[1]
	xy := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

	set := make([][]bool, row)
	for k := 0; k < row; k++ {
		set[k] = make([]bool, col)
	}

	if board[i][j] == 'M' {
		board[i][j] = 'X'
	} else if board[i][j] == 'E' {
		q := [][]int{{i, j}}
		set[i][j] = true
		for len(q) > 0 {
			top := q[0]
			q = q[1:]
			var count byte = 0
			for _, v := range xy {
				newI, newJ := top[0]+v[0], top[1]+v[1]
				if top[0]+v[0] < 0 || top[0]+v[0] >= row {
					continue
				}
				if top[1]+v[1] < 0 || top[1]+v[1] >= col {
					continue
				}
				if board[newI][newJ] == 'M' {
					count++
				}
			}
			if count > 0 {
				board[top[0]][top[1]] = 48 + count
			} else {
				board[top[0]][top[1]] = 'B'
				for _, v := range xy {
					newI, newJ := top[0]+v[0], top[1]+v[1]
					if top[0]+v[0] < 0 || top[0]+v[0] >= row {
						continue
					}
					if top[1]+v[1] < 0 || top[1]+v[1] >= col {
						continue
					}

					if board[newI][newJ] == 'E' && !set[newI][newJ] {
						set[newI][newJ] = true
						q = append(q, []int{newI, newJ})
					}
				}
			}
		}

	}

	return board
}

// solveSudoku
// 37. 解数独
// https://leetcode.cn/problems/sudoku-solver/
// TODO 待实现
func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	rows := make([]int, 9)
	cols := make([]int, 9)
	unit := make([]int, 9)
	// 初始化
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}

			v := 1 << int(board[i][j]-'1')
			unitI := i/3*3 + j/3
			rows[i] |= v
			cols[j] |= v
			unit[unitI] |= v
		}
	}

}
