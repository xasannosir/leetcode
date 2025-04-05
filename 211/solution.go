package solution

type WordDictionary struct {
	words []string
}

func Constructor() WordDictionary {
	return WordDictionary{
		words: []string{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.words = append(this.words, word)
}

func (this *WordDictionary) Search(word string) bool {
	for _, data := range this.words {
		equal := true
		if len(data) == len(word) {
			for i := 0; i < len(word); i++ {
				if word[i] != data[i] && word[i] != '.' {
					equal = false
					break
				}
			}

			if equal {
				return true
			}
		}
	}

	return false
}
