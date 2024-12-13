package solution

func minimumOperations(nums []int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		if nums[i]%3 != 0 {
			sum++
		}
	}

	return sum
}
