package solution

func rotate(nums []int, k int) {
	k = k % len(nums)

	for i := 0; i < k; i++ {
		temp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}
