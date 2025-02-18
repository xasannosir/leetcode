package solution

func smallestNumber(pattern string) string {
	pattern += "I"

	mn := 1
	res := make([]byte, len(pattern))

	for i, v := range pattern {
		if v == 'I' {
			res[i] = 48 + byte(mn)
			mn++

			if i > 0 && pattern[i-1] == 'D' {
				for j := i - 1; j >= 0 && pattern[j] == 'D'; j-- {
					res[j] = 48 + byte(mn)
					mn++
				}
			}
		}
	}

	return string(res)
}
