package require

import (
	"slices"
	"testing"
)

func EqualSlices(t *testing.T, caseIndex int, actual, result []int) {
	if slices.Equal(result, actual) {
		return
	}
	t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v",
		caseIndex, result, actual)
}
