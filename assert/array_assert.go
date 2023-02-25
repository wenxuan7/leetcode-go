package assert

import (
	"math"
	"testing"
)

// VerifyFloat64 校验值是否相等
func VerifyFloat64(t *testing.T, caseIndex int, result, actual, p float64) {
	if math.Dim(result, actual) > p {
		t.Fatalf("结果与实际值不同, caseIndex: %d, result: %f, actual: %f",
			caseIndex, result, actual)
	}
}

// Verify 校验值是否相等
func Verify[C comparable](t *testing.T, caseIndex int, result, actual C) {
	if result != actual {
		t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v",
			caseIndex, result, actual)
	}
}

// VerifyArr 校验数组全部值是否相等
func VerifyArr[C comparable](t *testing.T, caseIndex int, result, actual []C) {
	if len(result) != len(actual) {
		t.Fatalf("len must be equal, result: %v, resultLen: %d, actual: %v, actualLen: %d",
			result, len(result), actual, len(actual))
	}

	for i := range actual {
		if result[i] != actual[i] {
			t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v, i: %d",
				caseIndex, result, actual, i)
		}
	}
}

// Verify2Arr 校验二维数组全部值是否相等
func Verify2Arr[C comparable](t *testing.T, caseIndex int, result, actual [][]C) {
	if len(result) != len(actual) {
		t.Fatalf("len must be equal, caseIndex: %d, result: %v, resultLen: %d, actual: %v, actualLen: %d",
			caseIndex, result, len(result), actual, len(actual))
	}

	for i := range actual {
		for j := range actual[i] {
			if actual[i][j] != result[i][j] {
				t.Fatalf("结果与实际值不同, caseIndex: %d, result: %v, actual: %v, i: %d",
					caseIndex, result, actual, i)
			}
		}
	}
}
