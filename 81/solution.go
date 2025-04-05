package solution

func search(nums []int, target int) bool {
	for _, num := range nums {
		if target == num {
			return true
		}
	}

	return false
}
