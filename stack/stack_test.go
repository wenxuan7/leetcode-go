package stack

import (
	"testing"
)

var (
	isValidData = []string{
		"",
		"{{{{[[[]]](())}}",
		"[}",
		"[]{{}}((())){([])}",
	}
	isValidActual = []bool{
		false,
		false,
		false,
		true,
	}
)

func TestIsValid(t *testing.T) {
	for i := range isValidData {
		result := isValid(isValidData[i])
		if result != isValidActual[i] {
			t.Fatalf("结果与实际不相符, caseIndex: %d, data: %s, result: %v, actual: %v",
				i, isValidData[i], result, isValidActual[i])
		}
	}
}

var (
	minStackOperate = [][]string{
		{"push", "push", "push", "getMin", "pop", "top", "getMin"},
	}
	minStackData = [][]int{
		{-2, 0, -3},
	}
	minStackActual = [][]int{
		{-3, 0, -2},
	}
)

func TestMinStack(t *testing.T) {
	ms := Constructor()
	for i := range minStackOperate {
		dataI, actualI := 0, 0
		for j := range minStackOperate[i] {
			switch minStackOperate[i][j] {
			case "push":
				ms.Push(minStackData[i][dataI])
				dataI++
			case "pop":
				ms.Pop()
			case "top":
				result := ms.Top()
				if result != minStackActual[i][actualI] {
					t.Fatalf("结果与实际不相符, caseIndex: %d, operate:%s, result: %d, actual: %d",
						i, minStackOperate[i][j], result, minStackActual[i][actualI])
				}
				actualI++
			case "getMin":
				result := ms.GetMin()
				if result != minStackActual[i][actualI] {
					t.Fatalf("结果与实际不相符, caseIndex: %d, operate:%s, result: %d, actual: %d",
						i, minStackOperate[i][j], result, minStackActual[i][actualI])
				}
				actualI++
			}
		}
	}

}
