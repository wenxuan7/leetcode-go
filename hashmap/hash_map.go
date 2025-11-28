package hashmap

func longestConsecutive(nums []int) int {
	visited := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		visited[v] = struct{}{}
	}

	ans, cnt := 0, 1
	ok := false
	for k := range visited {
		if _, ok = visited[k-1]; ok {
			continue
		}
		cnt = 1
		for i := 1; i < len(nums); i++ {
			if _, ok = visited[k+i]; ok {
				cnt++
				continue
			}
			break
		}
		ans = max(ans, cnt)
	}
	return ans
}
