package solution

func numEquivDominoPairs(dominoes [][]int) int {
	count := 0
	freq := make(map[[2]int]int)

	for _, domino := range dominoes {
		key := [2]int{domino[0], domino[1]}
		if domino[0] > domino[1] {
			key = [2]int{domino[1], domino[0]}
		}

		count += freq[key]

		freq[key]++
	}

	return count
}
