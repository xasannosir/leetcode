package solution

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	res := make([][]int, 0)
	pairs := make(map[int]int)

	for _, num := range nums1 {
		pairs[num[0]] += num[1]
	}

	for _, num := range nums2 {
		pairs[num[0]] += num[1]
	}

	for k, v := range pairs {
		res = append(res, []int{k, v})
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res)-i-1; j++ {
			if res[j][0] > res[j+1][0] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}

	return res
}
