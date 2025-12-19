package alg

// SegmentTree 线段树
type SegmentTree struct {
	n    int
	tree []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	st := &SegmentTree{
		n:    len(nums),
		tree: make([]int, 4*len(nums)),
	}
	st.init(1, 0, st.n-1, nums)
	return st
}

func (st *SegmentTree) init(node, l, r int, nums []int) {
	if l == r {
		st.tree[node] = nums[l]
		return
	}

	mid := l + (r-l)>>1
	st.init(node*2, l, mid, nums)
	st.init(node*2+1, mid+1, r, nums)

	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

func (st *SegmentTree) Update(index, value int) {
	st.update(1, 0, st.n-1, index, value)
}

func (st *SegmentTree) update(node, l, r, idx, val int) {
	if l == r {
		st.tree[node] = val
		return
	}

	mid := l + (r-l)>>1
	if idx <= mid {
		st.update(node*2, l, mid, idx, val)
	} else {
		st.update(node*2+1, mid+1, r, idx, val)
	}

	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

func (st *SegmentTree) RangeSum(ql, qr int) int {
	return st.rangeSum(1, 0, st.n-1, ql, qr)
}

func (st *SegmentTree) rangeSum(node, l, r, ql, qr int) int {
	// 完全包含
	if ql <= l && r <= qr {
		return st.tree[node]
	}

	// 完全不相交
	if r < ql || l > qr {
		return 0
	}

	mid := l + (r-l)>>1
	leftSum := st.rangeSum(node*2, l, mid, ql, qr)
	rightSum := st.rangeSum(node*2+1, mid+1, r, ql, qr)

	return leftSum + rightSum
}
