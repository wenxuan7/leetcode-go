package recursive

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	data := [][]int{
		{2, 3, 6, 7},
	}
	dataTarget := []int{
		7,
	}

	for i := range data {
		sum := combinationSum(data[i], dataTarget[i])
		fmt.Println(sum)
	}
}
