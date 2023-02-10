package stack

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	var (
		data = []string{
			"",
			"{{{{[[[]]](())}}",
			"[}",
			"[]{{}}((())){([])}",
		}
		actual = []bool{
			false,
			false,
			false,
			true,
		}
	)

	for i := range data {
		result := isValid(data[i])
		if result != actual[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, data: %s, result: %v, actual: %v",
				i, data[i], result, actual[i])
		}
	}
}

func TestMinStack(t *testing.T) {
	var (
		operate = [][]string{
			{"push", "push", "push", "getMin", "pop", "top", "getMin"},
		}
		data = [][]int{
			{-2, 0, -3},
		}
		actual = [][]int{
			{-3, 0, -2},
		}
	)

	ms := Constructor()
	for i := range operate {
		dataI, actualI := 0, 0
		for j := range operate[i] {
			switch operate[i][j] {
			case "push":
				ms.Push(data[i][dataI])
				dataI++
			case "pop":
				ms.Pop()
			case "top":
				result := ms.Top()
				if result != actual[i][actualI] {
					t.Fatalf("结果与实际不相符, caseIndex: %d, operate:%s, result: %d, actual: %d",
						i, operate[i][j], result, actual[i][actualI])
				}
				actualI++
			case "getMin":
				result := ms.GetMin()
				if result != actual[i][actualI] {
					t.Fatalf("结果与实际不相符, caseIndex: %d, operate:%s, result: %d, actual: %d",
						i, operate[i][j], result, actual[i][actualI])
				}
				actualI++
			}
		}
	}

}
