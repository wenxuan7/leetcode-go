package hashmap

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	data1 := []string{
		"anagram",
		"aaaaabb",
	}
	data2 := []string{
		"nagaram",
		"aaaaab",
	}
	actual := []bool{
		true,
		false,
	}

	for i := range data1 {
		result := isAnagram(data1[i], data2[i])
		if result != actual[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, result: %v, actual: %v",
				i, result, actual[i])
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	actual := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}

	result := groupAnagrams(data)
	t.Logf("result: %v, actual: %v\n",
		result, actual)
}
