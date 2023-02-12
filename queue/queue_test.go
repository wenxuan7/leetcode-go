package queue

import "testing"

func TestMyCircularDeque(t *testing.T) {
	operate := []string{
		"MyCircularDeque",
		"insertLast",
		"insertLast",
		"insertFront",
		"insertFront",
		"getRear",
		"isFull",
		"deleteLast",
		"insertFront",
		"getFront",
	}
	data := []int{
		3,
		1,
		2,
		3,
		4,
		4,
	}

	actualNum := []int{
		2,
		4,
	}
	actualBool := []bool{
		true,
		true,
		true,
		false,
		true,
		true,
		true,
	}
	dataI, actualNumI, actualBoolI := 0, 0, 0

	var dq MyCircularDeque
	for i := range operate {
		switch operate[i] {
		case "MyCircularDeque":
			dq = Constructor(data[dataI])
			dataI++
		case "insertFront":
			result := dq.InsertFront(data[dataI])
			dataI++
			if result != actualBool[actualBoolI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualBool[actualBoolI])
			}
			actualBoolI++
		case "insertLast":
			result := dq.InsertLast(data[dataI])
			dataI++
			if result != actualBool[actualBoolI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualBool[actualBoolI])
			}
			actualBoolI++
		case "deleteFront":
			result := dq.DeleteFront()
			if result != actualBool[actualBoolI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualBool[actualBoolI])
			}
			actualBoolI++
		case "deleteLast":
			result := dq.DeleteLast()
			if result != actualBool[actualBoolI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualBool[actualBoolI])
			}
			actualBoolI++
		case "getFront":
			result := dq.GetFront()
			if result != actualNum[actualNumI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualNum[actualNumI])
			}
			actualNumI++
		case "getRear":
			result := dq.GetRear()
			if result != actualNum[actualNumI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualNum[actualNumI])
			}
			actualNumI++
		case "isEmpty":
			result := dq.IsEmpty()
			if result != actualBool[actualBoolI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualBool[actualBoolI])
			}
			actualBoolI++
		case "isFull":
			result := dq.IsFull()
			if result != actualBool[actualBoolI] {
				t.Fatalf("结果与实际不相符, i: %d, operate: %s, result: %v, actual: %v",
					i, operate[i], result, actualBool[actualBoolI])
			}
			actualBoolI++
		}
	}
}
