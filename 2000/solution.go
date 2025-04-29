package solution

func reversePrefix(word string, ch byte) string {
	reverse := func(s string) string {
		var swap string

		for i := len(s) - 1; i >= 0; i-- {
			swap += string(s[i])
		}

		return swap
	}

	findIndex := func(s string, symbol byte) int {
		for i := 0; i < len(s); i++ {
			if s[i] == symbol {
				return i
			}
		}

		return -1
	}

	index := findIndex(word, ch)

	if index == -1 {
		return word
	} else {
		return reverse(word[:index+1]) + word[index+1:]
	}
}
