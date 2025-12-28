package solution

func makeFancyString(s string) string {
	prev := s[0]
	prevCount := 1
	res := make([]byte, 0)
	res = append(res, s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			prevCount++
		} else {
			prev = s[i]
			prevCount = 1
		}

		if prevCount < 3 {
			res = append(res, s[i])
		}
	}

	return string(res)
}
