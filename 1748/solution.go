package solution

func sumOfUnique(nums []int) int {
	var sum int
	counter := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	for k, v := range counter {
		if v == 1 {
			sum += k
		}
	}

	return sum
}
