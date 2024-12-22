package solution

func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			return nums1[i]
		}
	}

	return 0
}
