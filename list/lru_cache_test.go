package list

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	op := [][]string{
		{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
	}
	data := [][][]int{
		{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
	}
	actual := [][]int{
		{1, -1, -1, 3, 4},
	}

	var lru LRUCache
	actualI := 0
	for i := range op {
		for j, s := range op[i] {
			switch s {
			case "LRUCache":
				lru = Constructor(data[i][j][0])
			case "get":
				result := lru.Get(data[i][j][0])
				assert.Verify(t, i, result, actual[i][actualI])
				actualI++
			case "put":
				lru.Put(data[i][j][0], data[i][j][1])
			}
		}
	}
}
