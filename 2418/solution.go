package solution

func sortPeople(names []string, heights []int) []string {
	for i := 0; i < len(heights)-1; i++ {
		flag := true
		for j := 0; j < len(heights)-i-1; j++ {
			if heights[j] < heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
				names[j], names[j+1] = names[j+1], names[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}

	return names
}
