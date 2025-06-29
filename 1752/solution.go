package solution

func check(nums []int) bool {
	sort := func(arr []int) {
		for i := 0; i < len(arr)-1; i++ {
			flag := true
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					flag = false
				}
			}

			if flag {
				break
			}
		}
	}

	isEqual := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	c := make([]int, len(nums))

	copy(c, nums)

	sort(c)

	for i := range nums {
		if isEqual(append(nums[i:], nums[:i]...), c) {
			return true
		}
	}

	return false
}
