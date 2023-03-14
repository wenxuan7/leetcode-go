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
