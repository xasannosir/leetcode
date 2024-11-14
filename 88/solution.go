package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; i <= m+n-1; i, j = i+1, j+1 {
		nums1[i] = nums2[j]
	}

	for i := 0; i < m+n; i++ {
		for j := i + 1; j < m+n; j++ {
			if nums1[i] > nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
}
