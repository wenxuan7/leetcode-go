package queue

import (
	. "leetcode-go/data"
	"math"
)

// 102. 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		level := make([]int, 0)

		levelCount := len(q)
		for i := 0; i < levelCount; i++ {
			level = append(level, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		q = q[levelCount:]
		ans = append(ans, level)
	}

	return ans
}

// minMutation
// 433. 最小基因变化
// https://leetcode.cn/problems/minimum-genetic-mutation/
func minMutation(startGene string, endGene string, bank []string) int {
	noExist := 0
	for i := 0; i < len(bank); i++ {
		if bank[i] != endGene {
			noExist++
		}
	}
	if noExist == len(bank) {
		return -1
	}

	q := make([]string, 0)
	q = append(q, startGene)

	ans := 1
	for len(q) > 0 {
		// 防止死循环
		if ans > len(bank) {
			return -1
		}
		levelCount := len(q)
		for i := 0; i < levelCount; i++ {
			if isModify(q[i], endGene) {
				return ans
			}

			for j := 0; j < len(bank); j++ {
				if isModify(q[i], bank[j]) {
					q = append(q, bank[j])
				}
			}
		}

		if len(q) == levelCount {
			return -1
		}
		q = q[levelCount:]
		ans++
	}

	return -1
}

func isModify(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	bs1, bs2 := []byte(str1), []byte(str2)
	count := 0
	for i := 0; i < len(bs1); i++ {
		if bs1[i] != bs2[i] {
			count++
		}

		if count > 1 {
			return false
		}
	}

	return count == 1
}

// largestValues
// 515. 在每个树行中找最大值
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		levelMax := math.MinInt

		levelCount := len(q)
		for i := 0; i < levelCount; i++ {
			if q[i].Val > levelMax {
				levelMax = q[i].Val
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}

		q = q[levelCount:]
		ans = append(ans, levelMax)
	}

	return ans
}

// ladderLength
// 127. 单词接龙
// https://leetcode.cn/problems/word-ladder/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]int, len(beginWord)*(len(wordList)+1))
	graph := make([][]int, 0)
	addM := func(str *string) int {
		id, ok := m[*str]
		if !ok {
			id = len(m)
			m[*str] = id
			graph = append(graph, []int{})
		}
		return id
	}

	generateGraph := func(str *string) int {
		id1 := addM(str)

		bs := []byte(*str)
		for i, v := range bs {
			bs[i] = '*'
			s := string(bs)
			id2 := addM(&s)
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			bs[i] = v
		}

		return id1
	}

	for i := range wordList {
		generateGraph(&wordList[i])
	}

	startId := generateGraph(&beginWord)
	endId, ok := m[endWord]
	if !ok {
		return 0
	}

	const inf = math.MaxInt64
	distStart := make([]int, len(m))
	for i := range distStart {
		distStart[i] = inf
	}
	distStart[startId] = 0

	distEnd := make([]int, len(m))
	for i := range distEnd {
		distEnd[i] = inf
	}
	distEnd[endId] = 0

	qStart := []int{startId}
	qEnd := []int{endId}
	for len(qStart) > 0 && len(qEnd) > 0 {
		l := len(qStart)

		for i := 0; i < l; i++ {
			v := qStart[i]
			if distEnd[v] < inf {
				return (distEnd[v]+distStart[v])/2 + 1
			}
			for j := range graph[v] {
				if distStart[graph[v][j]] == inf {
					distStart[graph[v][j]] = distStart[v] + 1
					qStart = append(qStart, graph[v][j])
				}
			}
		}
		qStart = qStart[l:]

		l = len(qEnd)
		for i := 0; i < l; i++ {
			v := qEnd[i]
			if distStart[v] < inf {
				return (distEnd[v]+distStart[v])/2 + 1
			}
			for j := range graph[v] {
				if distEnd[graph[v][j]] == inf {
					distEnd[graph[v][j]] = distEnd[v] + 1
					qEnd = append(qEnd, graph[v][j])
				}
			}
		}
		qEnd = qEnd[l:]
	}

	return 0
}
