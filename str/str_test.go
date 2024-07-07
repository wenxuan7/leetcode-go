package str

import (
	"github.com/leetcode/require"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	data := []int{
		3749,
		58,
	}
	actual := []string{
		"MMMDCCXLIX",
		"LVIII",
	}
	for i, num := range data {
		roman := intToRoman(num)
		require.EqualN(t, i, actual[i], roman)
	}
}
