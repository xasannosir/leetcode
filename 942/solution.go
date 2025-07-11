package solution

func diStringMatch(s string) []int {
	n := len(s)
	perm := make([]int, n+1)
	start := 0
	end := n

	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			perm[i] = start
			start++
		} else {
			perm[i] = end
			end--
		}
	}

	if s[n-1] == 'I' {
		perm[n] = start
	} else {
		perm[n] = end
	}

	return perm
}
