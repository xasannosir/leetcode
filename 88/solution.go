package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; i <= m+n-1; i, j = i+1, j+1 {
		nums1[i] = nums2[j]
	}

	for i := 0; i < len(nums1)-1; i++ {
		flag := true
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}
}
