package solution

func searchRange(nums []int, target int) []int {
	firstInd, secondInd := -1, -1
	firstHave, secondHave := false, false

	for i, val := range nums {
		if val == target && !firstHave {
			firstInd = i
			firstHave = true
		} else if val == target && firstHave {
			secondInd = i
			secondHave = true
		}
	}

	if firstHave && !secondHave {
		secondInd = firstInd
	}

	return []int{firstInd, secondInd}
}
