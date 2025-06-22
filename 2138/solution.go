package solution

func divideString(s string, k int, fill byte) []string {
	res := make([]string, 0)

	i := 0
	g := ""

	for ; i < len(s); i++ {
		if len(g) == k {
			res = append(res, g)
			g = ""
		}

		g += string(s[i])
	}

	for len(g) != k {
		g += string(fill)
	}
	
	res = append(res, g)

	return res
}
