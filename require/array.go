package require

import (
	"cmp"
	"slices"
	"testing"
)

func Equal1D[T cmp.Ordered](t *testing.T, caseIdx int, actual, result []T) {
	if slices.Equal(result, actual) {
		return
	}
	t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v",
		caseIdx, result, actual)
}

func Equal2D[T cmp.Ordered](t *testing.T, caseIdx int, actual, result [][]T) {
	if len(result) != len(actual) {
		t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v",
			caseIdx, result, actual)
		return
	}
	for i := range actual {
		if !slices.Equal(result[i], actual[i]) {
			t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v",
				caseIdx, result, actual)
		}
	}
}
