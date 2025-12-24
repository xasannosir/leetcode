package solution

func minimizedStringLength(s string) int {
	counter := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		counter[s[i]] = struct{}{}
	}

	return len(counter)
}
