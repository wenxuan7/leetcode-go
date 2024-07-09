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
