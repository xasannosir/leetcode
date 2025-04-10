package solution

func canConstruct(ransomNote string, magazine string) bool {
	rCount := make(map[int]int)
	mCount := make(map[int]int)

	for i := 0; i < len(ransomNote); i++ {
		rCount[int(ransomNote[i])]++
	}

	for i := 0; i < len(magazine); i++ {
		mCount[int(magazine[i])]++
	}

	for k, v := range rCount {
		if count, ok := mCount[k]; !ok || v > count {
			return false
		}
	}

	return true
}
