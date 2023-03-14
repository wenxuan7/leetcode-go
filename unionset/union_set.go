package unionset

// findCircleNum
// 547. 省份数量
// https://leetcode.cn/problems/number-of-provinces/
func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 {
		return 0
	}

	parent := make([]int, len(isConnected))
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(start, end int) {
		parent[find(start)] = find(end)
	}

	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 0 {
				continue
			}
			union(i, j)
		}
	}

	ans := 0
	for i, v := range parent {
		if v == i {
			ans++
		}
	}
	return ans
}

// numIslands
// 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	us := NewUnionSet(grid)
	m, n := len(grid), len(grid[0])
	xy := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == '0' {
				continue
			}

			grid[i][j] = '0'

			for _, nums := range xy {
				newI, newJ := i+nums[0], j+nums[1]
				if newI >= 0 && newJ >= 0 &&
					newI < m &&
					newJ < n &&
					grid[newI][newJ] == '1' {
					us.Union(i*n+j, newI*n+newJ)
				}
			}

		}
	}

	return us.GetCount()
}

type UnionSet struct {
	nums  []int
	rank  []int
	count int
}

func NewUnionSet(grid [][]byte) UnionSet {
	m, n := len(grid), len(grid[0])
	us := UnionSet{nums: make([]int, m*n), rank: make([]int, m*n)}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}

			us.nums[i*n+j] = i*n + j
			us.count++
		}
	}

	return us
}

func (us *UnionSet) Find(x int) int {
	if us.nums[x] != x {
		us.nums[x] = us.Find(us.nums[x])
	}

	return us.nums[x]
}

func (us *UnionSet) Union(x, y int) {
	xv, yv := us.Find(x), us.Find(y)

	if xv != yv {
		if us.rank[xv] > us.rank[yv] {
			us.nums[yv] = xv
		} else if us.rank[xv] < us.rank[yv] {
			us.nums[xv] = yv
		} else {
			us.nums[yv] = xv
			us.rank[xv] += 1
		}

		us.count--
	}
}

func (us *UnionSet) GetCount() int {
	return us.count
}
