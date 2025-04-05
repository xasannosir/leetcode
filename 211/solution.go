package main

import "fmt"

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

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dat")
	obj.AddWord("mad")
	fmt.Println(obj.Search("pad") == false)
	fmt.Println(obj.Search("bad") == true)
	fmt.Println(obj.Search(".ad") == true)
	fmt.Println(obj.Search("b..") == true)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
