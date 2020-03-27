package problem0007

import (
	"math"
)

func reverse(x int) int {
	sign := 1
	res := 0

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		temp := x % 10
		res = res * 10 + temp
		x = x / 10
	}

	if res > math.MaxInt32 {
		res = 0
	}

	return sign * res
}
