package recursive

import (
	. "github.com/leetcode-go/data"
	"math"
	"sort"
	"strconv"
	"strings"
)

// generateParenthesis
// 22. 括号生成
// https://leetcode.cn/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	ans := make([]string, 0)
	var recursive func(bs []byte, start, count int)
	recursive = func(bs []byte, start, count int) {
		if count < 0 {
			return
		}

		if start == 2*n {
			if count == 0 {
				ans = append(ans, string(bs))
			}
			return
		}

		bs[start] = '('
		recursive(bs, start+1, count+1)

		bs[start] = ')'
		recursive(bs, start+1, count-1)
	}

	recursive(make([]byte, 2*n), 0, 0)

	return ans
}

// invertTree
// 226. 翻转二叉树
// https://leetcode.cn/problems/invert-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// isValidBST
// 98. 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return recursiveValidBST(root, math.MinInt64, math.MaxInt64)
}

func recursiveValidBST(tn *TreeNode, lower, upper int) bool {
	if tn == nil {
		return true
	}

	if !(tn.Val > lower && tn.Val < upper) {
		return false
	}

	return recursiveValidBST(tn.Left, lower, tn.Val) &&
		recursiveValidBST(tn.Right, tn.Val, upper)
}

// maxDepth
// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDep, rDep := 1+maxDepth(root.Left), 1+maxDepth(root.Right)

	if lDep > rDep {
		return lDep
	} else {
		return rDep
	}
}

// minDepth
// 111. 二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	lDep, rDep := math.MaxInt32, math.MaxInt32
	if root.Left != nil {
		lDep = 1 + minDepth(root.Left)
	}
	if root.Right != nil {
		rDep = 1 + minDepth(root.Right)
	}

	if lDep > rDep {
		return rDep
	} else {
		return lDep
	}
}

// Codec
// 297. 二叉树的序列化与反序列化
// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (cc *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var (
		q       []*TreeNode
		err     error
		builder = strings.Builder{}
	)

	q = append(q, root)
	for len(q) > 0 {
		top := q[0]

		str := "nil"
		if top != nil {
			str = strconv.Itoa(top.Val)
		}
		if _, err = builder.WriteString(str); err != nil {
			panic(err)
		}

		q = q[1:]
		if top != nil {
			q = append(q, top.Left)
			q = append(q, top.Right)
		}

		if len(q) > 0 {
			if _, err = builder.WriteString(","); err != nil {
				panic(err)
			}
		}
	}

	return builder.String()
}

// Deserializes your encoded data to tree.
func (cc *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	ss := strings.Split(data, ",")

	var (
		v   int
		err error
	)

	// 队列
	q := make([]*TreeNode, 0, len(ss))
	if v, err = strconv.Atoi(ss[0]); err != nil {
		panic(err)
	}

	root := &TreeNode{Val: v}
	q = append(q, root)
	for j := 1; j < len(ss); {
		if ss[j] != "nil" {
			if v, err = strconv.Atoi(ss[j]); err != nil {
				panic(err)
			}

			left := &TreeNode{Val: v}
			q[0].Left = left
			q = append(q, left)
		}
		j++

		if ss[j] != "nil" {
			if v, err = strconv.Atoi(ss[j]); err != nil {
				panic(err)
			}

			right := &TreeNode{Val: v}
			q[0].Right = right
			q = append(q, right)
		}
		j++
		// 出队列
		q = q[1:]
	}

	return root
}

// lowestCommonAncestor
// 236. 二叉树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}
}

// combine
// 77. 组合
// https://leetcode.cn/problems/combinations/
func combine(n int, k int) [][]int {
	if k > n {
		return [][]int{}
	}

	ans := make([][]int, 0)
	nums := make([]int, 0, k)

	var dfs func(start, k int, nums []int)
	dfs = func(start, count int, nums []int) {
		if 0 == count {
			dst := make([]int, k)
			copy(dst, nums)
			ans = append(ans, dst)
			return
		}

		for i := start; i <= n-count+1; i++ {
			nums = append(nums, i)
			dfs(i+1, count-1, nums)
			nums = nums[:len(nums)-1]
		}
	}

	dfs(1, k, nums)
	return ans
}

// 46. 全排列
// https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	ansCount := 1
	for i := 1; i <= len(nums); i++ {
		ansCount *= i
	}

	ans := make([][]int, 0, ansCount)
	has := make([]bool, len(nums))
	var recursive func(count int, temp []int)
	recursive = func(count int, temp []int) {
		if count == len(nums) {
			dst := make([]int, len(nums))
			copy(dst, temp)
			ans = append(ans, dst)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !has[i] {
				has[i] = true
				temp = append(temp, nums[i])
				recursive(count+1, temp)
				temp = temp[:len(temp)-1]
				has[i] = false
			}
		}
	}

	recursive(0, make([]int, 0, len(nums)))
	return ans
}

// permuteUnique
// 47. 全排列 II
// https://leetcode.cn/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	sort.Ints(nums)

	ansCount := 1
	for i := 1; i <= len(nums); i++ {
		ansCount *= i
	}

	ans := make([][]int, 0, ansCount)
	has := make([]bool, len(nums))
	var recursive func(count int, temp []int)
	recursive = func(count int, temp []int) {
		if count == len(nums) {
			dst := make([]int, len(nums))
			copy(dst, temp)
			ans = append(ans, dst)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 重复的数 且前一个没有
			if i != 0 && nums[i] == nums[i-1] && !has[i-1] {
				continue
			}

			if !has[i] {
				has[i] = true
				temp = append(temp, nums[i])
				recursive(count+1, temp)
				temp = temp[:len(temp)-1]
				has[i] = false
			}
		}
	}

	recursive(0, make([]int, 0, len(nums)))
	return ans
}

// myPow
// 50. Pow(x, n)
// https://leetcode.cn/problems/powx-n/
func myPow(x float64, n int) float64 {
	var recursive func(x float64, n int) float64
	recursive = func(x float64, n int) float64 {
		if n == 1 {
			return x
		}

		pow := myPow(x, n/2)
		if n%2 == 0 {
			return pow * pow
		}
		return pow * pow * x
	}

	if n == 0 {
		return 1
	} else if n > 0 {
		return recursive(x, n)
	} else {
		return 1 / recursive(x, -n)
	}
}

// subsets
// 78. 子集
// https://leetcode.cn/problems/subsets/solutions/
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	var ans [][]int
	var recursive func(i int, temp []int)
	recursive = func(i int, temp []int) {
		if i == len(nums) {
			dst := make([]int, len(temp))
			copy(dst, temp)
			ans = append(ans, dst)
			return
		}

		temp = append(temp, nums[i])
		recursive(i+1, temp)
		temp = temp[:len(temp)-1]
		recursive(i+1, temp)
	}

	recursive(0, make([]int, 0, len(nums)))
	return ans
}
