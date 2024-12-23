package solution

func searchRange(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = append(result, i)
		}
	}

	if len(result) == 0 {
		return []int{-1, -1}
	} else if len(result) == 1 {
		return []int{result[0], result[0]}
	} else if len(result) > 2 {
		return []int{result[0], result[len(result)-1]}
	}

	return result
}
