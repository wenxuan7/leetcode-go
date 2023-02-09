package stack

import "container/list"

// isValid
// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	if s == "" {
		return false
	}

	ls := list.New()

	chs := []byte(s)
	for i := range chs {
		ch := chs[i]
		if ch == '{' || ch == '[' || ch == '(' {
			ls.PushBack(ch)
			continue
		}

		if ls.Len() == 0 {
			return false
		}

		back := ls.Back()
		topCh := back.Value.(byte)
		if (topCh == '{' && ch == '}') ||
			(topCh == '[' && ch == ']') ||
			(topCh == '(' && ch == ')') {
			ls.Remove(back)
		} else {
			return false
		}
	}

	return ls.Len() == 0
}

// MinStack
// 155. 最小栈
// https://leetcode.cn/problems/min-stack/
type MinStack struct {
	ls  *list.List
	min int
}

func Constructor() MinStack {
	return MinStack{ls: list.New()}
}

func (ms *MinStack) Push(val int) {
	if ms.ls.Len() == 0 {
		ms.ls.PushBack(0)
		ms.min = val
		return
	}

	sub := val - ms.min
	ms.ls.PushBack(sub)
	if sub < 0 {
		ms.min = val
	}
}

func (ms *MinStack) Pop() {
	if ms.ls.Len() == 0 {
		return
	}

	back := ms.ls.Back()
	v := back.Value.(int)
	ms.ls.Remove(back)

	if v <= 0 {
		ms.min = ms.min - v
	}
}

func (ms *MinStack) Top() int {
	if ms.ls.Len() == 0 {
		return 0
	}

	back := ms.ls.Back().Value
	v := back.(int)
	if v >= 0 {
		return v + ms.min
	}
	return ms.min
}

func (ms *MinStack) GetMin() int {
	return ms.min
}
