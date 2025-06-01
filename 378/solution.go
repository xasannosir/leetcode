package solution

func kthSmallest(matrix [][]int, k int) int {
	sort := func(arr []int) {
		for i := 0; i < len(arr); i++ {
			flag := false
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					flag = true
				}
			}

			if !flag {
				break
			}
		}
	}

	nums := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			nums = append(nums, matrix[i][j])
		}
	}

	sort(nums)

	return nums[k-1]
}
