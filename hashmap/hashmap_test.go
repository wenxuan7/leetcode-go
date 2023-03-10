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

func TestTwoSum(t *testing.T) {
	var (
		data = [][]int{
			{},
			{1, 2, 3, 4, 5, 6},
			{1, 2, 3, 4, 5, 6},
		}
		target = []int{
			1,
			9,
			12,
		}
		actual = [][]int{
			{},
			{3, 4},
			{},
		}
	)

	for i := range data {
		result := twoSum(data[i], target[i])
		verifyArr(t, i, result, actual[i])
	}
}

// verifyArr 校验数组全部值是否相等
func verifyArr(t *testing.T, caseIndex int, result []int, actual []int) {
	if len(result) != len(actual) {
		t.Fatalf("len must be equal, result: %v, resultLen: %d, actual: %v, actualLen: %d",
			result, len(result), actual, len(actual))
	}

	for i := range actual {
		if result[i] != actual[i] {
			t.Fatalf("结果与实际值不同, caseIndex: %d, result: %d, actual: %d, i: %d",
				caseIndex, result, actual, i)
		}
	}
}
