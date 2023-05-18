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

func TestMyAtoi(t *testing.T) {
	s := []string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
	}
	actual := []int{
		42,
		-42,
		4193,
		0,
	}

	for i := range s {
		result := myAtoi(s[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestMaskPII(t *testing.T) {
	s := []string{
		"LeetCode@LeetCode.com",
		"AB@qq.com",
		"1(234)567-890",
		"86-(10)12345678",
	}
	actual := []string{
		"l*****e@leetcode.com",
		"a*****b@qq.com",
		"***-***-7890",
		"+**-***-***-5678",
	}

	for i := range s {
		result := maskPII(s[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestIsValid(t *testing.T) {
	str := []string{
		"aabcbc",
		"abcabcababcc",
		"abccba",
	}
	actual := []bool{
		true,
		true,
		false,
	}

	for i := range str {
		result := isValid(str[i])
		assert.Verify(t, i, result, actual[i])
	}
}

func TestGcdOfStrings(t *testing.T) {
	str1 := []string{
		"ABCABC",
		"ABABAB",
		"LEET",
	}
	str2 := []string{
		"ABC",
		"AB",
		"CODE",
	}
	actual := []string{
		"ABC",
		"AB",
		"",
	}
	for i := range str1 {
		result := gcdOfStrings(str1[i], str2[i])
		assert.Verify(t, i, result, actual[i])
	}
}
