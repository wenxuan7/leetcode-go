package tree

// Trie 208. 实现 Trie (前缀树)
// https://leetcode.cn/problems/implement-trie-prefix-tree/description/?envType=study-plan-v2&envId=top-interview-150
type Trie struct {
	children []*Trie
	end      bool
}

func Constructor() Trie {
	return Trie{children: make([]*Trie, 26), end: false}
}

func (t *Trie) Insert(word string) {
	pre, cur := t, t
	idx := 0
	for i := 0; i < len(word); i++ {
		idx = int(word[i] - 'a')
		if cur.children[idx] == nil {
			pre = cur
			cur = &Trie{children: make([]*Trie, 26), end: false}
			pre.children[idx] = cur
		} else {
			cur = cur.children[idx]
		}
	}
	cur.end = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	idx := 0
	for i := 0; i < len(word); i++ {
		idx = int(word[i]) - 'a'
		if cur == nil || cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return cur.end
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	idx := 0
	for i := 0; i < len(prefix); i++ {
		idx = int(prefix[i]) - 'a'
		if cur == nil || cur.children[idx] == nil {
			return false
		}
		cur = cur.children[idx]
	}
	return true
}

// WordDictionary 211. 添加与搜索单词 - 数据结构设计
// https://leetcode.cn/problems/design-add-and-search-words-data-structure/description/?envType=study-plan-v2&envId=top-interview-150
type WordDictionary struct {
	children []*WordDictionary
	end      bool
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{children: make([]*WordDictionary, 26), end: false}
}

func (wd *WordDictionary) AddWord(word string) {
	pre, cur := wd, wd
	idx, l := 0, len(word)
	for i := 0; i < l; i++ {
		idx = int(word[i] - 'a')
		if cur.children[idx] == nil {
			pre = cur
			cur = &WordDictionary{children: make([]*WordDictionary, 26)}
			pre.children[idx] = cur
		} else {
			cur = cur.children[idx]
		}
	}
	cur.end = true
}

func (wd *WordDictionary) Search(word string) bool {
	cur := wd
	idx := 0
	for i := 0; i < len(word); i++ {
		idx = int(word[i] - 'a')
		if cur == nil {
			return false
		}

		if word[i] != '.' {
			if cur.children[idx] == nil {
				return false
			}
			cur = cur.children[idx]
			continue
		}

		for j := 0; j < 26; j++ {
			if cur.children[j] == nil {
				continue
			}
			if cur.children[j].Search(word[i+1:]) {
				return true
			}
		}
		return false
	}
	return cur.end
}

// BoardLetter 212. 单词搜索 II
// https://leetcode.cn/problems/word-search-ii/submissions/545376905/?envType=study-plan-v2&envId=top-interview-150
type BoardLetter struct {
	children []*BoardLetter
	word     string
}

func (bl *BoardLetter) Insert(word string) {
	cur := bl
	idx := 0
	for i := 0; i < len(word); i++ {
		idx = int(word[i]) - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &BoardLetter{children: make([]*BoardLetter, 26)}
		}
		cur = cur.children[idx]
	}
	cur.word = word
}

func findWords(board [][]byte, words []string) []string {
	root := &BoardLetter{children: make([]*BoardLetter, 26)}
	for _, word := range words {
		root.Insert(word)
	}

	appeared := make(map[string]bool, len(words))
	var dfs func(i, j int, root *BoardLetter)
	dfs = func(i, j int, root *BoardLetter) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
			return
		}
		if board[i][j] == '.' {
			return
		}

		idx := board[i][j] - 'a'
		if root.children[idx] == nil {
			return
		}
		if root.children[idx].word != "" && !appeared[root.children[idx].word] {
			appeared[root.children[idx].word] = true
		}

		board[i][j] = '.'
		dfs(i+1, j, root.children[idx])
		dfs(i, j+1, root.children[idx])
		dfs(i, j-1, root.children[idx])
		dfs(i-1, j, root.children[idx])
		board[i][j] = 'a' + idx
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, root)
		}
	}

	ret := make([]string, 0, len(words))
	for word := range appeared {
		ret = append(ret, word)
	}
	return ret
}
