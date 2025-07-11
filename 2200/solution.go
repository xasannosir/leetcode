package solution

func findKDistantIndices(nums []int, key int, k int) []int {
	keyIndices := make([]int, 0)
	distinctIndices := make([]int, 0)
	distinctChecker := make(map[int]struct{})

	abs := func(num int) int {
		if num < 0 {
			return -num
		}

		return num
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			keyIndices = append(keyIndices, i)
		}
	}

	for i := 0; i < len(keyIndices); i++ {
		for j := 0; j < len(nums); j++ {
			if _, ok := distinctChecker[j]; abs(j-keyIndices[i]) <= k && !ok {
				distinctIndices = append(distinctIndices, j)
				distinctChecker[j] = struct{}{}
			}
		}
	}

	return distinctIndices
}
