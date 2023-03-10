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
