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

	for i := 0; i < len(res)-1; i++ {
		flag := true
		for j := 0; j < len(res)-i-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}

	return res
}
