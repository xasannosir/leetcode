package solution

func intersection(nums [][]int) []int {
	inArray := func(val int, array []int) bool {
		for i := 0; i < len(array); i++ {
			if array[i] == val {
				return true
			}
		}

		return false
	}

	var res []int

	for _, v := range nums[0] {
		isHave := true
		for i := 1; i < len(nums); i++ {
			if !inArray(v, nums[i]) {
				isHave = false
				break
			}
		}

		if isHave {
			res = append(res, v)
		}
	}

	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}

	return res
}
