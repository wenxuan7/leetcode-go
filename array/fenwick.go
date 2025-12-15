package array

// Fenwick 离散化树状数组
type Fenwick []int

func (fw Fenwick) Add(i, num int) {
	for j := i; j < len(fw); j += j & -j {
		fw[j] += num
	}
}

func (fw Fenwick) Sum(i int) int {
	if i < 1 {
		return 0
	}
	sum := 0
	for j := i; j > 0; j -= j & -j {
		sum += fw[j]
	}
	return sum
}

func (fw Fenwick) RangeSum(l, r int) int {
	if l > r {
		return 0
	}
	if l < 1 {
		l = 1
	}
	if r >= len(fw) {
		r = len(fw) - 1
	}
	return fw.Sum(r) - fw.Sum(l-1)
}
