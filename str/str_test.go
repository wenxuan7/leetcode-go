package str

import (
	"github.com/leetcode-go/assert"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	s := []string{
		"Hello",
		"here",
		"LOVELY",
	}
	actual := []string{
		"hello",
		"here",
		"lovely",
	}

	for i := range s {
		result := toLowerCase(s[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestLengthOfLastWord(t *testing.T) {
	s := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}
	actual := []int{
		5,
		4,
		6,
	}

	for i := range s {
		result := lengthOfLastWord(s[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestNumJewelsInStones(t *testing.T) {
	jewels := []string{
		"aA",
		"z",
	}
	stones := []string{
		"aAAbbbb",
		"ZZ",
	}
	actual := []int{
		3,
		0,
	}

	for i := range jewels {
		result := numJewelsInStones(jewels[i], stones[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestFirstUniqChar(t *testing.T) {
	s := []string{
		"leetcode",
		"loveleetcode",
		"aabb",
		"dddccdbba",
	}
	actual := []int{
		0,
		2,
		-1,
		8,
	}

	for i := range s {
		result := firstUniqChar(s[i])
		assert.Verify(t, i, result, actual[i])
	}
}
