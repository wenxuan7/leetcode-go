package stack

import (
	"container/list"
	"strconv"
)

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

// evalRPN
// 150. 逆波兰表达式求值
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	top := -1
	for _, token := range tokens {
		switch token {
		case "+":
			stack[top-1] = stack[top] + stack[top-1]
			top--
		case "-":
			stack[top-1] = stack[top-1] - stack[top]
			top--
		case "*":
			stack[top-1] = stack[top] * stack[top-1]
			top--
		case "/":
			stack[top-1] = stack[top-1] / stack[top]
			top--
		default:
			atoi, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}

			top++
			stack[top] = atoi
		}
	}
	return stack[top]
}

// 224. 基本计算器
// https://leetcode.cn/problems/basic-calculator/
// 去括号初始前面补+
// -(1+(4+5+2)-3)+(6+8)
func calculate(s string) int {
	ans := 0
	sign := 1
	op := make([]int, len(s))
	op[0] = 1
	top := 0

	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
			continue
		case '+':
			sign = 1 * op[top]
			i++
		case '-':
			sign = -1 * op[top]
			i++
		case '(':
			top++
			op[top] = sign
			i++
		case ')':
			top--
			i++
		default:
			num := 0
			for ; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += num * sign
		}
	}

	return ans
}
