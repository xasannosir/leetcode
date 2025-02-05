package solution

func check(nums []int) bool {
	sort := func(arr []int) {
		for i := 0; i < len(arr); i++ {
			for j := i; j < len(arr); j++ {
				if arr[i] > arr[j] {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
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

	for i := range nums {
		c[i] = nums[i]
	}

	sort(c)

	for i := range nums {
		if isEqual(append(nums[i:], nums[:i]...), c) {
			return true
		}
	}

	return false
}
