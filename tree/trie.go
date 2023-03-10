package tree

// Trie
// 208. 实现 Trie (前缀树)
// https://leetcode.cn/problems/implement-trie-prefix-tree/
type Trie struct {
	letter [26]*Trie
	isWord bool
}

func Constructor() Trie {
	return Trie{letter: [26]*Trie{}}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	if t == nil {
		*t = Trie{letter: [26]*Trie{}}
	}

	curr, key := t, byte(0)
	for i := range word {
		key = word[i] - 'a'
		if curr.letter[key] == nil {
			curr.letter[key] = &Trie{letter: [26]*Trie{}}
		}
		curr = curr.letter[key]
	}

	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	curr, key := t, byte(0)
	for i := range word {
		if curr == nil {
			return false
		}

		key = word[i] - 'a'
		if curr.letter[key] == nil {
			return false
		}
		curr = curr.letter[key]
	}

	return curr.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	curr, key := t, byte(0)
	for i := range prefix {
		if curr == nil {
			return false
		}

		key = prefix[i] - 'a'
		if curr.letter[key] == nil {
			return false
		}
		curr = curr.letter[key]
	}

	return true
}

// findWords
// 212. 单词搜索 II
// https://leetcode.cn/problems/word-search-ii/
func findWords(board [][]byte, words []string) []string {
	var ans []string
	m := len(board)
	n := len(board[0])
	// build tree
	root := NewTrieNode()
	for i := 0; i < len(words); i++ {
		insert(root, &words[i])
	}
	// stop backtracking early
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, root, &ans)
		}
	}
	return ans
}

func dfs(board [][]byte, i int, j int, node *TrieNode, ans *[]string) {
	ch := board[i][j] - 'a'
	m := len(board)
	n := len(board[0])
	if node.next[ch] == nil {
		return
	}
	next := node.next[ch]
	if next.word != nil {
		*ans = append(*ans, *(next.word))
		next.word = nil
	}
	// 还存在叶子节点
	if next.count > 0 {
		board[i][j] = 0 // mark
		i += 1          // down
		if i < m && board[i][j] != 0 {
			dfs(board, i, j, next, ans)
		}
		i -= 1
		j += 1 // right
		if j < n && board[i][j] != 0 {
			dfs(board, i, j, next, ans)
		}
		j -= 1
		i -= 1 // up
		if i >= 0 && board[i][j] != 0 {
			dfs(board, i, j, next, ans)
		}
		i += 1
		j -= 1 // left
		if j >= 0 && board[i][j] != 0 {
			dfs(board, i, j, next, ans)
		}
		j += 1
		board[i][j] = 'a' + ch
	}
	// 已经完成了当前节点为根的所有树枝的遍历，由于叶子节点的count为0，可以依靠递归摘除
	if next.count == 0 {
		node.count--
		node.next[ch] = nil
	}
}

type TrieNode struct {
	next  []*TrieNode
	count int
	word  *string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		next:  make([]*TrieNode, 26),
		count: 0,
		word:  nil,
	}
}

func insert(node *TrieNode, word *string) {
	for _, ch := range *word {
		if node.next[ch-'a'] == nil {
			node.next[ch-'a'] = NewTrieNode()
			node.count += 1 // child 数量
		}
		node = node.next[ch-'a']
	}
	node.word = word
}
