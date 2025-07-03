package solution

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	answer1, answer2 := 0, 0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				answer1++
				break
			}
		}
	}

	for i := 0; i < len(nums2); i++ {
		for j := 0; j < len(nums1); j++ {
			if nums2[i] == nums1[j] {
				answer2++
				break
			}
		}
	}

	return []int{answer1, answer2}
}
