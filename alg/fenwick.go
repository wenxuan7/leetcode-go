package alg

// Fenwick 离散树状数组
type Fenwick []int

func (fw Fenwick) Add(k, v int) {
	if k < 1 || k > len(fw)-1 {
		return
	}
	for i := k; i < len(fw); i += i & -i {
		fw[i] += v
	}
}

func (fw Fenwick) Sum(i int) int {
	if i < 1 || i > len(fw)-1 {
		return 0
	}
	sum := 0
	for j := i; j > 0; j -= j & -j {
		sum += fw[j]
	}
	return sum
}

func (fw Fenwick) RangeSum(l, r int) int {
	if l > r ||
		l < 1 || l > len(fw)-1 ||
		r < 1 || r > len(fw)-1 {
		return 0
	}
	return fw.Sum(r) - fw.Sum(l-1)
}
