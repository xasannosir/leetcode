package solution

func toLowerCase(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if 'A' <= runes[i] && runes[i] <= 'Z' {
			runes[i] = runes[i] + 32
		}
	}

	return string(runes)
}
