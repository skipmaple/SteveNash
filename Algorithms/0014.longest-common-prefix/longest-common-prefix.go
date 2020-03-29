package problem0014

func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	for i, c := range short {
		for _, str := range strs {
			if str[i] != byte(c) {
				return short[:i]
			}
		}
	}

	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, str := range strs {
		if len(res) > len(str) {
			res = str
		}
	}

	return res
}
