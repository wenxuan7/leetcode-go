package array

import "testing"

func TestFenwick(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fw := make(Fenwick, len(nums)+1)
	for i, num := range nums {
		fw.Add(i+1, num)
	}
	for i := 0; i < len(nums); i++ {
		t.Log(fw.Sum(i + 1))
	}
}
