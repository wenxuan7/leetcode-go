package tool

import "golang.org/x/exp/constraints"

func Max[C constraints.Ordered](x, y C) C {
	if x > y {
		return x
	}
	return y
}

func Min[C constraints.Ordered](x, y C) C {
	if x < y {
		return x
	}
	return y
}
