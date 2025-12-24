package solution

func minimumBoxes(apple []int, capacity []int) int {
	sort := func(nums []int) {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums)-i-1; j++ {
				if nums[j] < nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	}

	sort(capacity)

	total := 0

	for i := 0; i < len(apple); i++ {
		total += apple[i]
	}

	count := 0

	for i := 0; i < len(capacity); i++ {
		total -= capacity[i]
		count++

		if total <= 0 {
			return count
		}
	}

	return 0
}
