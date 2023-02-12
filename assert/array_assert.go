package assert

import "testing"

// Verify 校验值是否相等
func Verify(t *testing.T, caseIndex int, result int, actual int) {
	if result != actual {
		t.Fatalf("结果与实际值不同, caseIndex: %d, result: %d, actual: %d",
			caseIndex, result, actual)
	}
}

// VerifyArr 校验数组全部值是否相等
func VerifyArr(t *testing.T, caseIndex int, result []int, actual []int) {
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

// VerifySecArr 校验二维数组全部值是否相等
func VerifySecArr(t *testing.T, caseIndex int, result [][]int, actual [][]int) {
	if len(result) != len(actual) {
		t.Fatalf("len must be equal, result: %v, resultLen: %d, actual: %v, actualLen: %d",
			result, len(result), actual, len(actual))
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
