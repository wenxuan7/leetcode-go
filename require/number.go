package require

import (
	"cmp"
	"testing"
)

func EqualN[T cmp.Ordered](t *testing.T, i int, result, actual T) {
	if result != actual {
		t.Fatalf("结果错误, case: %d, result: %v, actual: %v",
			i, result, actual)
	}
}
