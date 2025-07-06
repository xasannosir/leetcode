package solution

func uniqueMorseRepresentations(words []string) int {
	pairs := map[byte]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	set := make(map[string]struct{})

	for _, word := range words {
		var morse string

		for i := 0; i < len(word); i++ {
			morse += pairs[word[i]]
		}

		set[morse] = struct{}{}
	}

	return len(set)
}
