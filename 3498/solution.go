package solution

func reverseDegree(s string) int {
	res := 0

	for i, val := range s {
		res += (123 - int(rune(val))) * (i + 1)
	}

	return res
}
