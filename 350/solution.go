package solution

func intersect(nums1 []int, nums2 []int) []int {
	freq := make(map[int]int)
	result := make([]int, 0)

	for _, num := range nums1 {
		freq[num]++
	}

	for _, num := range nums2 {
		if count, ok := freq[num]; ok && count > 0 {
			result = append(result, num)
			freq[num]--
		}
	}

	return result
}
