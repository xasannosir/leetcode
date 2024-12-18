package solution

func decodeMessage(key string, message string) string {
	opener := make(map[rune]rune)
	var start = 'a'

	for i := 0; i < len(key); i++ {
		if key[i] != ' ' {
			if _, ok := opener[rune(key[i])]; !ok {
				opener[rune(key[i])] = start
				start++
			}
		}
	}
	opener[' '] = ' '

	var result string

	for i := 0; i < len(message); i++ {
		result += string(opener[rune(message[i])])
	}

	return result
}
