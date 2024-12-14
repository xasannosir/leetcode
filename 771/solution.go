package solution

func numJewelsInStones(jewels string, stones string) int {
	var count int

	for _, jewel := range jewels {
		for _, stone := range stones {
			if jewel == stone {
				count++
			}
		}
	}

	return count
}
