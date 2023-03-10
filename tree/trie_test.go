package tree

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestTrie(t *testing.T) {
	operate := [][]string{
		{"insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "search", "search", "search", "search", "search", "search", "search", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith"},
	}
	data := [][]string{
		{"app", "apple", "beer", "add", "jam", "rental", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam"},
	}
	actual := [][]string{
		{"null", "null", "null", "null", "null", "null", "false", "true", "false", "false", "false", "false", "false", "true", "true", "false", "true", "true", "false", "false", "false", "true", "true", "true"},
	}

	for i := range operate {
		trie := Constructor()
		result, hasResult := false, false
		for j := range operate[i] {
			hasResult = false
			switch operate[i][j] {
			case "insert":
				trie.Insert(data[i][j])
			case "search":
				result = trie.Search(data[i][j])
				hasResult = true
			case "startsWith":
				result = trie.StartsWith(data[i][j])
				hasResult = true
			}

			// 是否需要返回值判断
			if !hasResult {
				continue
			}

			if result {
				assert.Verify(t, i, "true", actual[i][j])
			} else {
				assert.Verify(t, i, "false", actual[i][j])
			}
		}
	}
}

func TestFindWords(t *testing.T) {
	board := [][][]byte{
		{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		{{'a', 'b'}, {'c', 'd'}},
	}
	words := [][]string{
		{"oath", "pea", "eat", "rain"},
		{"abcb"},
	}
	actual := [][]string{
		{"oath", "eat"},
		{},
	}

	for i := range board {
		result := findWords(board[i], words[i])
		assert.VerifyArr(t, i, result, actual[i])
	}
}
