package queue

import (
	. "github.com/leetcode-go/data"
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
	strToId := make(map[string]int, len(beginWord)*(len(wordList)+1))
	graph := make([][]int, 0)
	generateStrToId := func(str *string) int {
		id, ok := strToId[*str]
		if !ok {
			id = len(strToId)
			strToId[*str] = id
			graph = append(graph, []int{})
		}
		return id
	}

	generateGraph := func(str *string) int {
		id1 := generateStrToId(str)

		bs := []byte(*str)
		for i, v := range bs {
			bs[i] = '*'
			s := string(bs)
			id2 := generateStrToId(&s)
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			bs[i] = v
		}

		return id1
	}

	// 建立双向图
	for i := range wordList {
		generateGraph(&wordList[i])
	}

	startId := generateGraph(&beginWord)
	endId, ok := strToId[endWord]
	if !ok {
		return 0
	}

	// 存储已搜索的id
	const inf = math.MaxInt64
	distStart := make([]int, len(strToId))
	for i := range distStart {
		distStart[i] = inf
	}
	distStart[startId] = 0

	distEnd := make([]int, len(strToId))
	for i := range distEnd {
		distEnd[i] = inf
	}
	distEnd[endId] = 0

	qStart := []int{startId}
	qEnd := []int{endId}
	// 双向广度优先搜索
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

// findLadders
// 126. 单词接龙 II
// https://leetcode.cn/problems/word-ladder-ii/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ans := make([][]string, 0)
	graph := make(map[string][]string, len(wordList))

	// 检索出beginWord的下一层节点
	graph[beginWord] = make([]string, 0)
	for i := range wordList {
		if wordList[i] == beginWord {
			continue
		}
		graph[wordList[i]] = make([]string, 0)
		if isModify(wordList[i], beginWord) {
			graph[beginWord] = append(graph[beginWord], wordList[i])
			graph[wordList[i]] = append(graph[wordList[i]], beginWord)
		}
	}
	// 检索wordList里每一个的下一层节点
	for i := 0; i < len(wordList); i++ {
		if beginWord == wordList[i] {
			continue
		}
		for j := i + 1; j < len(wordList); j++ {
			if beginWord == wordList[j] {
				continue
			}
			if isModify(wordList[i], wordList[j]) {
				graph[wordList[i]] = append(graph[wordList[i]], wordList[j])
				graph[wordList[j]] = append(graph[wordList[j]], wordList[i])
			}
		}
	}

	strToLevel := make(map[string]int, len(wordList))
	strToLevel[beginWord] = 0
	q := []string{beginWord}
	searched := false
	for len(q) > 0 {
		top := q[0]
		q = q[1:]

		for _, s := range graph[top] {
			if _, ok := strToLevel[s]; !ok {
				strToLevel[s] = strToLevel[top] + 1
				if s == endWord {
					searched = true
					break
				}
				q = append(q, s)
			}
		}

		if searched {
			break
		}
	}

	if !searched {
		return ans
	}

	endToBeginPath := make([]string, 0, strToLevel[endWord]+1)
	endToBeginPath = append(endToBeginPath, endWord)
	var dfs func(curr string)
	dfs = func(curr string) {
		if curr == beginWord {
			miniPath := make([]string, 0, len(endToBeginPath))
			for i := len(endToBeginPath) - 1; i >= 0; i-- {
				miniPath = append(miniPath, endToBeginPath[i])
			}
			ans = append(ans, miniPath)
			return
		}

		for _, s := range graph[curr] {
			if strToLevel[s]+1 == strToLevel[curr] {
				endToBeginPath = append(endToBeginPath, s)
				dfs(s)
				endToBeginPath = endToBeginPath[:len(endToBeginPath)-1]
			}
		}
	}

	dfs(endWord)

	return ans
}
