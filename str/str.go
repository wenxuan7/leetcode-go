package str

var iToR = [][]string{
	{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	{"M", "MM", "MMM", "", "", "", "", "", ""},
}

// intToRoman 12. 整数转罗马数字
// https://leetcode.cn/problems/integer-to-roman/submissions/544790798/?envType=study-plan-v2&envId=top-interview-150
func intToRoman(num int) string {
	tmp := 0
	ret := make([]byte, 16)
	r := len(ret) - 1
	for i := 0; num > 0; i++ {
		tmp = num % 10
		num /= 10
		if tmp == 0 {
			continue
		}
		s := iToR[i][tmp-1]
		for j := len(s) - 1; j >= 0; j-- {
			ret[r] = s[j]
			r--
		}
	}
	return string(ret[r+1:])
}
