package problem0013

func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s)-1; i >= 0; i-- {
		sign := 1
		if m[s[i]] < last {
			sign = -1
		}
		res += sign * m[s[i]]
		last = m[s[i]]
	}

	return res
}
