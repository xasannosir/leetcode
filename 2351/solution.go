package solution

func repeatedCharacter(s string) byte {
	counter := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
		if counter[s[i]] == 2 {
			return s[i]
		}
	}

	return ' '
}
