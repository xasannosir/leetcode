package solution

func kWeakestRows(mat [][]int, k int) []int {
	nums := make([]int, 0)
	idx := make([]int, 0)

	for i := 0; i < len(mat); i++ {
		total := 0
		for j := 0; j < len(mat[i]); j++ {
			total += mat[i][j]
		}
		nums = append(nums, total)
		idx = append(idx, i)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				idx[j], idx[j+1] = idx[j+1], idx[j]
			}
		}
	}

	return idx[:k]
}
