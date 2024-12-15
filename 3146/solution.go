package solution

func findPermutationDifference(s string, t string) int {
	var res int

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				if i > j {
					res += i - j
				} else {
					res += j - i
				}
			}
		}
	}

	return res
}
