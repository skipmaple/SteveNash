package problem0038

func countAndSay(n int) string {
	buf := []byte{'1'}

	for i := 1; i < n; i++ {
		buf = say(buf)
	}

	return string(buf)
}

func say(buf []byte) []byte {
	res := make([]byte, 0, len(buf)*2)
	i, j := 0, 1
	for i < len(buf) {
		for j < len(buf) && buf[i] == buf[j] {
			j++
		}

		res = append(res, byte(j-i+'0'), buf[i])
		i = j
	}

	return res
}
