package solution

func removeStars(s string) string {
	res := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			count := 0

			j := i
			for ; j < len(s); j++ {
				if s[j] == '*' {
					count++
				} else {
					break
				}
			}

			res = res[:len(res)-count]

			i = j - 1
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}
