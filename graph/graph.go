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

var cross = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var rice = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, -1},
	{-1, 1},
	{-1, -1},
	{1, 1},
}

// solve 130. 被围绕的区域
// https://leetcode.cn/problems/surrounded-regions/?envType=study-plan-v2&envId=top-interview-150
func solve(board [][]byte) {
	if len(board) < 2 {
		return
	}

	rs, cs := len(board), len(board[0])
	for c, b := range board[0] {
		if b != 'O' {
			continue
		}
		resetA(board, 0, c)
	}
	for c, b := range board[rs-1] {
		if b != 'O' {
			continue
		}
		resetA(board, rs-1, c)
	}
	for r := range board {
		if r == 0 || r == rs-1 {
			continue
		}
		if board[r][0] == 'O' {
			resetA(board, r, 0)
		}
		if board[r][cs-1] == 'O' {
			resetA(board, r, cs-1)
		}
	}

	for r, row := range board {
		for c := range row {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
			if board[r][c] == 'X' {
				continue
			}
			board[r][c] = 'O'
		}
	}
}

func resetA(board [][]byte, r, c int) {
	if r < 0 ||
		c < 0 ||
		r >= len(board) ||
		c >= len(board[0]) {
		return
	}
	if board[r][c] != 'O' {
		return
	}
	board[r][c] = 'A'
	resetA(board, r-1, c)
	resetA(board, r+1, c)
	resetA(board, r, c-1)
	resetA(board, r, c+1)
}
