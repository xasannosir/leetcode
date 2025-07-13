package solution

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)

	reverse := func(nums []int) []int {
		res := make([]int, len(nums))

		for i, j := len(nums)-1, 0; i >= 0; i, j = i-1, j+1 {
			res[j] = nums[i]
		}

		return res
	}

	for i := 0; i < n; i++ {
		image[i] = reverse(image[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if image[i][j] == 1 {
				image[i][j] = 0
			} else {
				image[i][j] = 1
			}
		}
	}

	return image
}
