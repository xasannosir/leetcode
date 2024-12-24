package solution

func intersection(nums1 []int, nums2 []int) []int {
	var result []int

	notIn := func(arr []int, num int) bool {
		for i := 0; i < len(arr); i++ {
			if num == arr[i] {
				return false
			}
		}

		return true
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] && notIn(result, nums1[i]) {
				result = append(result, nums1[i])
			}
		}
	}

	return result
}
