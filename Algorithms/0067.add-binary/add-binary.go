package problem0067

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	maxLen := len(a)

	tA := transA(a, maxLen)
	tB := transA(b, maxLen)

	t := add(tA, tB)
	return transS(t)
}

func transS(t []int) string {
	res := make([]byte, len(t))
	for i,c := range t {
		res[i] = byte(c + '0')
	}
	return string(res)
}

func add(a []int, b []int) []int {
	carry := 0
	l := len(a)
	res := make([]int, l+1)
	for i := l-1; i >= 0; i-- {
		carry += a[i] + b[i]
		res[i+1] = carry % 2
		carry /= 2
	}

	res[0] = carry

	for i, v := range res{
		if v != 0 {
			return res[i:]
		}
	}

	return []int{0}
}

func transA(s string, maxLen int) []int {
	res := make([]int, maxLen)
	for i, ch := range s {
		res[maxLen - len(s) + i] = int(ch - '0')
	}

	return res
}
