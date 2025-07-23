package solution

func addedInteger(nums1 []int, nums2 []int) int {
	sort := func(arr []int) {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	sort(nums1)
	sort(nums2)

	return nums2[0] - nums1[0]
}
