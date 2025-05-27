package solution

func findDifference(nums1 []int, nums2 []int) [][]int {
	isHave := func(num int, arr []int) bool {
		for i := 0; i < len(arr); i++ {
			if arr[i] == num {
				return true
			}
		}

		return false
	}

	dist1 := make([]int, 0, len(nums1))
	dist2 := make([]int, 0, len(nums2))

	for _, num := range nums1 {
		if !isHave(num, nums2) && !isHave(num, dist1) {
			dist1 = append(dist1, num)
		}
	}

	for _, num := range nums2 {
		if !isHave(num, nums1) && !isHave(num, dist2) {
			dist2 = append(dist2, num)
		}
	}

	return [][]int{dist1, dist2}
}
